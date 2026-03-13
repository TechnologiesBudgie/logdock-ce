package fs

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"logdock/internal/storage"
)

// Store persists logs as append-only NDJSON partitions (one file per day).
//
// Concurrency (BUG-001 fix): switched from sync.Mutex to sync.RWMutex.
// Append holds a write lock; readDay/GetMetrics hold a read lock so concurrent
// dashboard reads no longer race with writes.
//
// File-handle cache (BUG-005 fix): instead of open+close on every Append, we
// keep a single buffered writer open for the active day partition, rotated only
// on day roll-over. This removes the dominant syscall cost at high ingest rates.
type Store struct {
	root       string
	mu         sync.RWMutex // BUG-001: was sync.Mutex
	total      int64
	errorCount int64
	minutely   []time.Time

	// BUG-005: cached write path.
	activeDay  string
	activeFile *os.File
	activeBuf  *bufio.Writer
}

// New initialises the data directory and returns a ready Store.
func New(root string) (*Store, error) {
	if err := os.MkdirAll(root, 0o750); err != nil {
		return nil, err
	}
	s := &Store{root: root, minutely: make([]time.Time, 0, 1024)}
	s.loadStats()
	return s, nil
}

func (s *Store) loadStats() {
	b, err := os.ReadFile(filepath.Join(s.root, "stats.json"))
	if err == nil {
		var stats struct {
			Total      int64 `json:"total"`
			ErrorCount int64 `json:"errorCount"`
		}
		if json.Unmarshal(b, &stats) == nil {
			s.total = stats.Total
			s.errorCount = stats.ErrorCount
		}
	}
}

func (s *Store) saveStats() {
	stats := struct {
		Total      int64 `json:"total"`
		ErrorCount int64 `json:"errorCount"`
	}{Total: s.total, ErrorCount: s.errorCount}
	b, _ := json.Marshal(stats)
	_ = os.WriteFile(filepath.Join(s.root, "stats.json"), b, 0o640)
}

// Append adds a record to the day-partition matching its UTC timestamp.
func (s *Store) Append(r storage.LogRecord) error {
	s.mu.Lock() // exclusive write lock
	defer s.mu.Unlock()

	day := r.Timestamp.UTC().Format("2006-01-02")
	if err := s.ensureWriterLocked(day); err != nil {
		return err
	}
	if err := json.NewEncoder(s.activeBuf).Encode(r); err != nil {
		return err
	}
	// Flush so concurrent readers see complete lines immediately.
	if err := s.activeBuf.Flush(); err != nil {
		return err
	}
	s.total++
	if strings.EqualFold(r.Level, "error") || strings.EqualFold(r.Level, "fatal") {
		s.errorCount++
	}
	s.minutely = append(s.minutely, time.Now().UTC())
	s.trimMinutelyLocked()
	s.saveStats()
	return nil
}

// ensureWriterLocked opens (or keeps) the buffered writer for the given day.
// Caller MUST hold s.mu exclusively.
func (s *Store) ensureWriterLocked(day string) error {
	if s.activeDay == day && s.activeFile != nil {
		return nil
	}
	if s.activeFile != nil {
		_ = s.activeBuf.Flush()
		_ = s.activeFile.Close()
		s.activeFile = nil
		s.activeBuf = nil
	}
	dir := filepath.Join(s.root, day)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return err
	}
	f, err := os.OpenFile(filepath.Join(dir, "logs.ndjson"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o640)
	if err != nil {
		return err
	}
	s.activeDay = day
	s.activeFile = f
	s.activeBuf = bufio.NewWriterSize(f, 64*1024)
	return nil
}

// Search performs a case-insensitive message filter over a single day.
func (s *Store) Search(day, term string, limit int) ([]storage.LogRecord, error) {
	if limit <= 0 {
		limit = 200
	}
	recs, err := s.readDay(day, 100_000)
	if err != nil {
		return nil, err
	}
	res := make([]storage.LogRecord, 0, limit)
	needle := strings.ToLower(term)
	for _, r := range recs {
		if needle == "" || strings.Contains(strings.ToLower(r.Message), needle) {
			res = append(res, r)
		}
		if len(res) >= limit {
			break
		}
	}
	return res, nil
}

// Query applies all filters across every day partition in [From, To].
// BUG-007 fix: the original code only examined the single day derived from
// filter.From, silently missing all records in subsequent days of a range.
func (s *Store) Query(filter storage.QueryFilter) (storage.QueryResult, error) {
	from := filter.From
	to := filter.To
	if from.IsZero() {
		from = time.Now().UTC().Truncate(24 * time.Hour)
	}
	if to.IsZero() {
		to = time.Now().UTC()
	}
	if filter.Limit <= 0 {
		filter.Limit = 200
	}

	// Enumerate every day in the closed interval [from, to].
	var days []string
	cur := from.UTC().Truncate(24 * time.Hour)
	end := to.UTC().Truncate(24 * time.Hour)
	for !cur.After(end) {
		days = append(days, cur.Format("2006-01-02"))
		cur = cur.Add(24 * time.Hour)
	}

	var all []storage.LogRecord
	for _, day := range days {
		recs, err := s.readDay(day, 100_000)
		if err != nil {
			return storage.QueryResult{}, err
		}
		all = append(all, recs...)
	}

	needle := strings.ToLower(filter.Term)
	streamSet := toSet(filter.Streams)
	levelSet := toSet(filter.Levels)

	res := make([]storage.LogRecord, 0, filter.Limit)
	for _, r := range all {
		if !filter.From.IsZero() && r.Timestamp.Before(filter.From) {
			continue
		}
		if !filter.To.IsZero() && r.Timestamp.After(filter.To) {
			continue
		}
		if needle != "" && !strings.Contains(strings.ToLower(r.Message), needle) {
			continue
		}
		if len(streamSet) > 0 && !streamSet[r.StreamID] {
			continue
		}
		if len(levelSet) > 0 && !levelSet[strings.ToLower(r.Level)] {
			continue
		}
		if len(filter.Fields) > 0 {
			ok := true
			for k, v := range filter.Fields {
				if r.Fields[k] != v {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}
		}
		res = append(res, r)
		if len(res) >= filter.Limit {
			break
		}
	}

	aggs := buildAggregations(filter.Aggregations, res)
	return storage.QueryResult{Records: res, Aggregations: aggs}, nil
}

// readDay loads all records from one day-partition under a shared read lock.
// BUG-001 fix: the original tail() had no lock at all.
func (s *Store) readDay(day string, limit int) ([]storage.LogRecord, error) {
	s.mu.RLock() // shared read lock – BUG-001 fix
	defer s.mu.RUnlock()

	base := filepath.Join(s.root, day, "logs.ndjson")
	var f *os.File
	var err error
	var isGz bool

	for _, ext := range []string{"", ".gz"} {
		f, err = os.Open(base + ext)
		if err == nil {
			isGz = ext == ".gz"
			break
		}
	}

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()

	var rdr io.Reader = f
	if isGz {
		gz, err := gzip.NewReader(f)
		if err != nil {
			return nil, err
		}
		defer gz.Close()
		rdr = gz
	}

	out := make([]storage.LogRecord, 0, 256)
	sc := bufio.NewScanner(rdr)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024) // up to 1 MB per line
	for sc.Scan() {
		var r storage.LogRecord
		if json.Unmarshal(sc.Bytes(), &r) == nil {
			out = append(out, r)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	if len(out) > limit {
		out = out[len(out)-limit:]
	}
	return out, nil
}

// GetMetrics returns real operational counters.
// BUG-008 fix: removed projectedARCHitRate and projectedThinPoolUsage —
// both were computed from modular arithmetic on the record count and bore
// no relation to actual system state.
func (s *Store) GetMetrics() (storage.Metrics, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.trimMinutelyLocked()
	diskBytes, partitionDays := s.measureDiskLocked()
	return storage.Metrics{
		TotalIngested:    s.total,
		ErrorCount:       s.errorCount,
		LastMinuteIngest: int64(len(s.minutely)),
		DiskBytes:        diskBytes,
		PartitionDays:    partitionDays,
	}, nil
}

// measureDiskLocked sums the sizes of all NDJSON partition files.
// Caller MUST hold s.mu.
func (s *Store) measureDiskLocked() (totalBytes int64, days int) {
	_ = filepath.WalkDir(s.root, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if strings.Contains(p, "logs.ndjson") {
			if fi, err2 := d.Info(); err2 == nil {
				totalBytes += fi.Size()
				if !d.IsDir() {
					days++
				}
			}
		}
		return nil
	})
	return totalBytes, days
}

// Close flushes buffered writes and releases the active file handle.
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.activeFile != nil {
		_ = s.activeBuf.Flush()
		err := s.activeFile.Close()
		s.activeFile = nil
		s.activeBuf = nil
		return err
	}
	return nil
}

func (s *Store) RunMaintenance(maxDays int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entries, err := os.ReadDir(s.root)
	if err != nil {
		return
	}

	now := time.Now().UTC()
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		day, err := time.Parse("2006-01-02", entry.Name())
		if err != nil {
			continue
		}

		ageDays := int(now.Sub(day).Hours() / 24)

		// Delete if past max retention
		if maxDays > 0 && ageDays > maxDays {
			_ = os.RemoveAll(filepath.Join(s.root, entry.Name()))
		}
	}
}

func (s *Store) trimMinutelyLocked() {
	cut := time.Now().UTC().Add(-1 * time.Minute)
	i := 0
	for i < len(s.minutely) && s.minutely[i].Before(cut) {
		i++
	}
	if i > 0 {
		s.minutely = append([]time.Time{}, s.minutely[i:]...)
	}
}

func toSet(ss []string) map[string]bool {
	if len(ss) == 0 {
		return nil
	}
	m := make(map[string]bool, len(ss))
	for _, s := range ss {
		m[strings.ToLower(s)] = true
	}
	return m
}

func buildAggregations(specs []string, records []storage.LogRecord) map[string]map[string]int64 {
	if len(specs) == 0 {
		return nil
	}
	result := make(map[string]map[string]int64, len(specs))
	for _, spec := range specs {
		parts := strings.SplitN(spec, ":", 2)
		if len(parts) != 2 || parts[0] != "count" {
			continue
		}
		field := parts[1]
		counts := make(map[string]int64)
		for _, r := range records {
			var val string
			switch field {
			case "level":
				val = strings.ToLower(r.Level)
			case "source":
				val = r.Source
			case "stream":
				val = r.StreamID
			default:
				val = r.Fields[field]
			}
			counts[val]++
		}
		result[field] = counts
	}
	return result
}

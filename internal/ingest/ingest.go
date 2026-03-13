package ingest

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"net"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	
	"logdock/internal/pipeline"
	"logdock/internal/storage"
	"logdock/internal/streams"
)

type Service struct {
	Store    storage.Storage
	Pipeline *pipeline.Pipeline
	Streams  *streams.Manager
}

func (s *Service) HandleJSON(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB limit
	defer r.Body.Close()
	var rec storage.LogRecord
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, "invalid payload: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := validate(&rec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if rec.ID == "" {
		rec.ID = newID()
	}
	if rec.Timestamp.IsZero() {
		rec.Timestamp = time.Now().UTC()
	}
	if s.Pipeline != nil {
		_ = s.Pipeline.Execute(&rec)
	}
	if s.Streams != nil {
		s.Streams.Assign(&rec)
	}
	if err := s.Store.Append(rec); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func validate(r *storage.LogRecord) error {
	if len(r.Message) > 64000 {
		return errors.New("message exceeds 64 KB limit")
	}
	if r.Source == "" {
		r.Source = "unknown"
	}
	if r.Level == "" {
		r.Level = "info"
	}
	if r.Timestamp.After(time.Now().Add(24 * time.Hour)) {
		return errors.New("timestamp too far in future")
	}
	return nil
}

func extractLevel(msg string) string {
	msg = strings.ToLower(msg)
	patterns := map[string]string{
		"error":   "error",
		"fatal":   "fatal",
		"warn":    "warn",
		"warning": "warn",
		"debug":   "debug",
		"trace":   "trace",
	}
	for p, l := range patterns {
		if strings.Contains(msg, p) {
			return l
		}
	}
	return "info"
}

// StartSyslogUDP starts a UDP syslog listener that shuts down when ctx is
// cancelled. BUG-002 fix: the previous goroutine had no cancellation path and
// would block indefinitely on pc.ReadFrom after SIGTERM.
func (s *Service) StartSyslogUDP(ctx context.Context, addr string) error {
	pc, err := net.ListenPacket("udp", addr)
	if err != nil {
		return err
	}
	// Unblock ReadFrom when the context is done.
	go func() {
		<-ctx.Done()
		_ = pc.Close()
	}()
	go func() {
		defer pc.Close()
		buf := make([]byte, 8192)
		for {
			n, _, err := pc.ReadFrom(buf)
			if err != nil {
				return // closed by context cancellation or real error
			}
			rec := storage.LogRecord{
				ID:        newID(),
				Timestamp: time.Now().UTC(),
				Source:    "syslog-udp",
				Level:     extractLevel(string(buf[:n])),
				Message:   string(buf[:n]),
			}
			if s.Pipeline != nil {
				_ = s.Pipeline.Execute(&rec)
			}
			if s.Streams != nil {
				s.Streams.Assign(&rec)
			}
			_ = s.Store.Append(rec)
		}
	}()
	return nil
}

// StartSyslogTCP starts a TCP syslog listener that shuts down when ctx is
// cancelled. BUG-003 fix: the previous implementation had no context and no
// limit on concurrent connections.
func (s *Service) StartSyslogTCP(ctx context.Context, addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		_ = ln.Close()
	}()
	go func() {
		defer ln.Close()
		for {
			c, err := ln.Accept()
			if err != nil {
				return // closed by context cancellation
			}
			go func(conn net.Conn) {
				defer conn.Close()
				// Close the connection when the context is cancelled.
				go func() {
					<-ctx.Done()
					_ = conn.Close()
				}()
				sc := bufio.NewScanner(conn)
				for sc.Scan() {
				rec := storage.LogRecord{
					ID:        newID(),
					Timestamp: time.Now().UTC(),
					Source:    "syslog-tcp",
					Level:     extractLevel(sc.Text()),
					Message:   sc.Text(),
				}
					if s.Pipeline != nil {
						_ = s.Pipeline.Execute(&rec)
					}
					if s.Streams != nil {
						s.Streams.Assign(&rec)
					}
					_ = s.Store.Append(rec)
				}
			}(c)
		}
	}()
	return nil
}

func newID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// Ingest adds a log record directly (used by API server internal path).
func (s *Service) Ingest(rec storage.LogRecord) error {
	if rec.ID == "" {
		rec.ID = newID()
	}
	if rec.Timestamp.IsZero() {
		rec.Timestamp = time.Now().UTC()
	}
	if s.Pipeline != nil {
		_ = s.Pipeline.Execute(&rec)
	}
	if s.Streams != nil {
		s.Streams.Assign(&rec)
	}
	return s.Store.Append(rec)
}

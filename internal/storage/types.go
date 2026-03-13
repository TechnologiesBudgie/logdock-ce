package storage

import (
	"time"
)

// LogRecord represents a normalized immutable event persisted by LogDock.
type LogRecord struct {
	ID        string            `json:"id"`
	Timestamp time.Time         `json:"timestamp"`
	Source    string            `json:"source"`
	Level     string            `json:"level"`
	Facility  string            `json:"facility,omitempty"`
	StreamID  string            `json:"stream_id,omitempty"`
	Message   string            `json:"message"`
	Fields    map[string]string `json:"fields,omitempty"`
}

// Metrics contains operational counters derived from ingested events.
type Metrics struct {
	TotalIngested    int64   `json:"totalIngested"`
	ErrorCount       int64   `json:"errorCount"`
	LastMinuteIngest int64   `json:"lastMinuteIngest"`
	DiskBytes        int64   `json:"diskBytes"`
	PartitionDays    int     `json:"partitionDays"`
	ArcHitRate       float64 `json:"arcHitRate"`
	ThinPoolUsage    float64 `json:"thinPoolUsage"`
}

// QueryFilter defines the parameters for searching and aggregating logs.
type QueryFilter struct {
	Term         string
	From         time.Time
	To           time.Time
	Streams      []string
	Levels       []string
	Fields       map[string]string
	Limit        int
	Aggregations []string // e.g., "count:level", "count:stream"
}

// QueryResult contains the matching records and aggregation summaries.
type QueryResult struct {
	Records      []LogRecord
	Aggregations map[string]map[string]int64 // e.g., "level": {"ERROR": 10, "INFO": 100}
}

// Storage defines the interface for persisting and querying logs.
type Storage interface {
	// Append adds a new log record to the store.
	Append(r LogRecord) error
	// Search queries the store for logs matching the term within the given day/range.
	// Deprecated: use Query instead.
	Search(day, term string, limit int) ([]LogRecord, error)
	// Query performs an advanced search with filtering and aggregation.
	Query(filter QueryFilter) (QueryResult, error)
	// GetMetrics returns current operational metrics.
	GetMetrics() (Metrics, error)
	// Close cleans up any resources used by the store.
	Close() error
}

package streams

import (
	"strings"

	"logdock/internal/storage"
)

// Stream defines a categorization for logs based on filter criteria.
type Stream struct {
	ID          string
	Title       string
	Description string
	Rules       []Rule
}

// Rule defines a single condition for a stream.
type Rule struct {
	Field string // e.g., "level", "source", "message"
	Value string
	Type  RuleType
}

type RuleType int

const (
	MatchExactly RuleType = iota
	Contains
	Regex
)

// Manager handles log categorization into streams.
type Manager struct {
	Streams []Stream
}

func (m *Manager) Assign(r *storage.LogRecord) {
	for _, s := range m.Streams {
		if m.match(s, r) {
			// In GrayLog, a log can belong to multiple streams. 
			// For simplicity, we set StreamID to the first match if it's empty,
			// or we could append to a fields list.
			if r.StreamID == "" {
				r.StreamID = s.ID
			}
		}
	}
}

func (m *Manager) match(s Stream, r *storage.LogRecord) bool {
	// BUG-016 fix: a stream with no rules is a catch-all ("All Logs" pattern).
	// The previous code returned false here, so the default "All Logs" stream
	// never matched anything and StreamID was never set.
	if len(s.Rules) == 0 {
		return true
	}
	// Default to AND matching for all rules
	for _, rule := range s.Rules {
		if !m.matchRule(rule, r) {
			return false
		}
	}
	return true
}

func (m *Manager) matchRule(rule Rule, r *storage.LogRecord) bool {
	var val string
	switch rule.Field {
	case "level":
		val = r.Level
	case "source":
		val = r.Source
	case "message":
		val = r.Message
	case "facility":
		val = r.Facility
	default:
		if r.Fields != nil {
			val = r.Fields[rule.Field]
		}
	}

	switch rule.Type {
	case MatchExactly:
		return val == rule.Value
	case Contains:
		return strings.Contains(val, rule.Value)
	default:
		return false
	}
}

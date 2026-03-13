// Package audit records security-relevant events to an in-memory ring-buffer.
// Events include HTTP requests, login attempts, user management, and config changes.
// Each entry captures the authenticated username when available.
package audit

import (
	"net/http"
	"sync"
	"time"
)

// EventType classifies audit events for filtering and alerting.
type EventType string

const (
	EventHTTP       EventType = "http"
	EventLogin      EventType = "auth.login"
	EventLogout     EventType = "auth.logout"
	EventLoginFail  EventType = "auth.login_fail"
	EventUserCreate EventType = "user.create"
	EventUserDelete EventType = "user.delete"
	EventKeyCreate  EventType = "key.create"
	EventKeyRevoke  EventType = "key.revoke"
	EventMFAChange  EventType = "mfa.change"
	EventConfigSet  EventType = "config.set"
	EventRuleCreate EventType = "pipeline.rule_create"
	EventRuleDelete EventType = "pipeline.rule_delete"
	EventAlertCreate EventType = "alert.create"
	EventAlertTrigger EventType = "alert.trigger"
	EventNodeAdd    EventType = "node.add"
	EventDataExport EventType = "data.export"
)

// Entry is one audited event.
type Entry struct {
	ID         string    `json:"id"`
	Time       time.Time `json:"time"`
	Type       EventType `json:"type"`
	Method     string    `json:"method,omitempty"`
	Path       string    `json:"path,omitempty"`
	RemoteAddr string    `json:"remote_addr"`
	Status     int       `json:"status,omitempty"`
	Username   string    `json:"username,omitempty"`
	Detail     string    `json:"detail,omitempty"`
	Success    bool      `json:"success"`
}

const maxEntries = 5000

// Service captures and stores audit entries.
type Service struct {
	mu      sync.Mutex
	entries []Entry
	counter uint64
}

// New returns a ready audit Service.
func New() *Service { return &Service{} }

// Log records an arbitrary event.
func (s *Service) Log(e Entry) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
	e.ID = time.Now().UTC().Format("20060102150405") + idSuffix(s.counter)
	if e.Time.IsZero() {
		e.Time = time.Now().UTC()
	}
	s.entries = append(s.entries, e)
	if len(s.entries) > maxEntries {
		s.entries = s.entries[len(s.entries)-maxEntries:]
	}
}

// Middleware wraps an http.Handler and records each request with optional username extraction.
func (s *Service) Middleware(usernameFromReq func(*http.Request) string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, code: http.StatusOK}
		next.ServeHTTP(rw, r)

		username := ""
		if usernameFromReq != nil {
			username = usernameFromReq(r)
		}

		s.Log(Entry{
			Type:       EventHTTP,
			Method:     r.Method,
			Path:       r.URL.Path,
			RemoteAddr: r.RemoteAddr,
			Status:     rw.code,
			Username:   username,
			Success:    rw.code < 400,
		})
	})
}

// List returns the last `limit` audit entries (newest first).
func (s *Service) List(limit int) []Entry {
	s.mu.Lock()
	defer s.mu.Unlock()
	if limit <= 0 || limit > len(s.entries) {
		limit = len(s.entries)
	}
	start := len(s.entries) - limit
	out := make([]Entry, limit)
	copy(out, s.entries[start:])
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}

// Stats returns a summary of recent audit activity.
func (s *Service) Stats() map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	types := map[EventType]int{}
	users := map[string]int{}
	failures := 0
	for _, e := range s.entries {
		types[e.Type]++
		if e.Username != "" {
			users[e.Username]++
		}
		if !e.Success {
			failures++
		}
	}
	return map[string]any{
		"total":    len(s.entries),
		"failures": failures,
		"by_type":  types,
		"by_user":  users,
	}
}

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (r *responseWriter) WriteHeader(code int) {
	r.code = code
	r.ResponseWriter.WriteHeader(code)
}

func idSuffix(n uint64) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 4)
	for i := range b {
		b[i] = chars[n%36]
		n /= 36
	}
	return string(b)
}

// Package sse provides a Server-Sent Events hub for real-time log streaming.
// Clients subscribe to streams and receive log records as JSON events.
package sse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"logdock/internal/storage"
)

// client is a connected SSE subscriber.
type client struct {
	ch       chan storage.LogRecord
	streamID string // "" = all streams
	level    string // "" = all levels
	term     string // substring filter
}

// Hub manages SSE clients and broadcasts log records.
type Hub struct {
	mu      sync.RWMutex
	clients map[*client]struct{}
}

// New returns a ready Hub.
func New() *Hub {
	return &Hub{clients: make(map[*client]struct{})}
}

// Publish sends a record to all matching subscribers (non-blocking).
func (h *Hub) Publish(rec storage.LogRecord) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for c := range h.clients {
		// Stream filter.
		if c.streamID != "" && rec.StreamID != c.streamID {
			continue
		}
		// Level filter.
		if c.level != "" && rec.Level != c.level {
			continue
		}
		// Term filter.
		if c.term != "" && !containsFold(rec.Message, c.term) && !containsFold(rec.Source, c.term) {
			continue
		}
		// Non-blocking send — drop if client is slow.
		select {
		case c.ch <- rec:
		default:
		}
	}
}

// ServeHTTP implements the SSE endpoint. Query params:
//   stream=<id>   filter to a specific stream
//   level=<level> filter by level
//   q=<term>      substring filter
func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no") // disable nginx buffering

	c := &client{
		ch:       make(chan storage.LogRecord, 64),
		streamID: r.URL.Query().Get("stream"),
		level:    r.URL.Query().Get("level"),
		term:     r.URL.Query().Get("q"),
	}

	h.mu.Lock()
	h.clients[c] = struct{}{}
	h.mu.Unlock()

	defer func() {
		h.mu.Lock()
		delete(h.clients, c)
		h.mu.Unlock()
		close(c.ch)
	}()

	// Send a connect event so the client knows the stream is live.
	fmt.Fprintf(w, "event: connected\ndata: {\"ts\":\"%s\"}\n\n", time.Now().UTC().Format(time.RFC3339))
	flusher.Flush()

	// Heartbeat ticker to keep the connection alive through proxies.
	heartbeat := time.NewTicker(15 * time.Second)
	defer heartbeat.Stop()

	for {
		select {
		case <-r.Context().Done():
			return
		case <-heartbeat.C:
			fmt.Fprintf(w, ": heartbeat %s\n\n", time.Now().UTC().Format(time.RFC3339))
			flusher.Flush()
		case rec, ok := <-c.ch:
			if !ok {
				return
			}
			b, _ := json.Marshal(rec)
			fmt.Fprintf(w, "event: log\ndata: %s\n\n", b)
			flusher.Flush()
		}
	}
}

// ClientCount returns the number of active SSE connections.
func (h *Hub) ClientCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.clients)
}

func containsFold(s, sub string) bool {
	if sub == "" {
		return true
	}
	sl := len(sub)
	for i := 0; i <= len(s)-sl; i++ {
		if equalFold(s[i:i+sl], sub) {
			return true
		}
	}
	return false
}

func equalFold(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		ca, cb := a[i], b[i]
		if ca >= 'A' && ca <= 'Z' {
			ca += 32
		}
		if cb >= 'A' && cb <= 'Z' {
			cb += 32
		}
		if ca != cb {
			return false
		}
	}
	return true
}

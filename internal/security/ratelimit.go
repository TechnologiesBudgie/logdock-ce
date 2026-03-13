package security

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	windowDuration = time.Minute
	maxRequests    = 240
	cleanupEvery   = 5 * time.Minute
	staleCutoff    = 10 * time.Minute
)

// BruteForceGuard applies IP-based sliding-window rate limiting.
//
// BUG-004 fix: the original map was never pruned — every unique IP address
// encountered during the process lifetime would accumulate indefinitely.
// We now run a background cleanup goroutine (started via Start) that evicts
// entries that have had no traffic for staleCutoff. The max memory footprint
// is bounded by active unique IPs × (windowDuration × requestsPerSecond).
type BruteForceGuard struct {
	mu   sync.Mutex
	hits map[string][]time.Time
}

// NewBruteForceGuard returns a new guard. Call Start to begin periodic cleanup.
func NewBruteForceGuard() *BruteForceGuard {
	return &BruteForceGuard{hits: make(map[string][]time.Time)}
}

// Start launches a background goroutine that prunes stale IP entries.
// It exits when ctx is cancelled.
func (g *BruteForceGuard) Start(ctx context.Context) {
	go func() {
		t := time.NewTicker(cleanupEvery)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				g.prune()
			}
		}
	}()
}

// prune removes entries that have had no requests in staleCutoff.
func (g *BruteForceGuard) prune() {
	g.mu.Lock()
	defer g.mu.Unlock()
	cut := time.Now().UTC().Add(-staleCutoff)
	for ip, times := range g.hits {
		// Keep entry only if there is at least one recent hit.
		recent := false
		for _, t := range times {
			if t.After(cut) {
				recent = true
				break
			}
		}
		if !recent {
			delete(g.hits, ip)
		}
	}
}

// Middleware enforces per-IP rate limiting.
func (g *BruteForceGuard) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1 MB DoS guard

		h, _, _ := net.SplitHostPort(r.RemoteAddr)
		if h == "" {
			h = r.RemoteAddr
		}
		if !g.allow(h) {
			http.Error(w, "rate limited", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (g *BruteForceGuard) allow(ip string) bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	now := time.Now().UTC()
	cut := now.Add(-windowDuration)

	cur := g.hits[ip]
	// Slide the window: drop hits older than windowDuration.
	filtered := cur[:0]
	for _, t := range cur {
		if t.After(cut) {
			filtered = append(filtered, t)
		}
	}
	if len(filtered) >= maxRequests {
		g.hits[ip] = filtered
		return false
	}
	g.hits[ip] = append(filtered, now)
	return true
}

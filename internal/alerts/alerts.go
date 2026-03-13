package alerts

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"logdock/internal/storage"
)

// Rule configures threshold or regex alerting.
type Rule struct {
	Name      string `json:"name"`
	Condition string `json:"condition"` // e.g., "level=ERROR > 10"
	Target    string `json:"target"`
	StreamID  string `json:"stream_id,omitempty"`
}

// Event is a triggered alert instance.
type Event struct {
	Rule      string    `json:"rule"`
	Severity  string    `json:"severity"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// Service stores rules/history and runs background evaluation.
type Service struct {
	mu      sync.Mutex
	rules   []Rule
	history []Event
	store   storage.Storage
}

func New(st storage.Storage) *Service {
	return &Service{
		store: st,
		rules: []Rule{{Name: "error-spike", Condition: "level=ERROR > 10", Target: "webhook"}},
	}
}

func (s *Service) Start(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.evaluate()
		}
	}
}

func (s *Service) evaluate() {
	rules := s.Rules()
	for _, r := range rules {
		// Very simplified evaluation logic
		// In GrayLog, this would query the store for the last N minutes
		now := time.Now().UTC()
		filter := storage.QueryFilter{
			From:  now.Add(-5 * time.Minute),
			To:    now,
			Limit: 1000,
		}
		if r.StreamID != "" {
			filter.Streams = []string{r.StreamID}
		}

		res, err := s.store.Query(filter)
		if err != nil {
			log.Printf("Alert evaluation error: %v", err)
			continue
		}

		// Basic threshold check: if level=ERROR > 10 (hardcoded for demo)
		count := 0
		for _, rec := range res.Records {
			if rec.Level == "ERROR" {
				count++
			}
		}

		if count > 10 {
			s.Push(Event{
				Rule:      r.Name,
				Severity:  "high",
				Message:   fmt.Sprintf("Threshold exceeded: count=%d", count),
				Timestamp: time.Now().UTC(),
			})
			log.Printf("[ALERT] triggered: %s", r.Name)
		}
	}
}
func (s *Service) Rules() []Rule {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Rule{}, s.rules...)
}
func (s *Service) History() []Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Event{}, s.history...)
}
func (s *Service) AddRule(r Rule) { s.mu.Lock(); defer s.mu.Unlock(); s.rules = append(s.rules, r) }
func (s *Service) Push(e Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.history = append(s.history, e)
	if len(s.history) > 200 {
		s.history = s.history[len(s.history)-200:]
	}
}

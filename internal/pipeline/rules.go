// Package pipeline provides a configurable log transformation engine.
// Rules are evaluated in order and can parse, drop, tag, redact, or
// re-route log records before they reach storage.
package pipeline

import (
	"encoding/json"
	"regexp"
	"strings"
	"sync"

	"logdock/internal/storage"
)

// ActionType describes what a rule does when it matches.
type ActionType string

const (
	ActionTag    ActionType = "tag"    // add a key=value field
	ActionDrop   ActionType = "drop"   // discard the record
	ActionRedact ActionType = "redact" // replace a field value with [REDACTED]
	ActionRoute  ActionType = "route"  // assign to a different stream
	ActionParse  ActionType = "parse"  // extract fields via named regex groups
	ActionLevel  ActionType = "level"  // override the log level
)

// Rule is a single pipeline transformation.
type Rule struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Enabled     bool       `json:"enabled"`
	// Match criteria (all non-empty fields must match, AND logic).
	MatchLevel  string     `json:"match_level,omitempty"`  // "error", "warn", etc.
	MatchSource string     `json:"match_source,omitempty"` // substring match
	MatchRegex  string     `json:"match_regex,omitempty"`  // applied to message
	// Action to take.
	Action      ActionType `json:"action"`
	// Action parameters.
	TagKey      string     `json:"tag_key,omitempty"`
	TagValue    string     `json:"tag_value,omitempty"`
	RedactField string     `json:"redact_field,omitempty"`
	RouteStream string     `json:"route_stream,omitempty"`
	ParseRegex  string     `json:"parse_regex,omitempty"` // named groups → fields
	LevelValue  string     `json:"level_value,omitempty"`

	compiled *regexp.Regexp
	parseRe  *regexp.Regexp
}

// compile pre-compiles regex fields. Must be called before Apply.
func (r *Rule) compile() error {
	if r.MatchRegex != "" {
		re, err := regexp.Compile(r.MatchRegex)
		if err != nil {
			return err
		}
		r.compiled = re
	}
	if r.ParseRegex != "" {
		re, err := regexp.Compile(r.ParseRegex)
		if err != nil {
			return err
		}
		r.parseRe = re
	}
	return nil
}

// matches returns true if the record satisfies all configured match criteria.
func (r *Rule) matches(rec storage.LogRecord) bool {
	if r.MatchLevel != "" && !strings.EqualFold(rec.Level, r.MatchLevel) {
		return false
	}
	if r.MatchSource != "" && !strings.Contains(strings.ToLower(rec.Source), strings.ToLower(r.MatchSource)) {
		return false
	}
	if r.compiled != nil && !r.compiled.MatchString(rec.Message) {
		return false
	}
	return true
}

// Engine holds the ordered list of pipeline rules and applies them to records.
type Engine struct {
	mu    sync.RWMutex
	rules []Rule
}

// New returns an Engine pre-loaded with sensible default rules.
func New() *Engine {
	e := &Engine{}
	defaults := []Rule{
		{
			ID: "default-redact-cc", Name: "Redact Credit Card Numbers",
			Enabled: true, MatchRegex: `\b\d{4}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}\b`,
			Action: ActionRedact, RedactField: "message",
		},
		{
			ID: "default-redact-token", Name: "Redact Bearer Tokens",
			Enabled: true, MatchRegex: `(?i)bearer\s+[a-z0-9\-_\.]+`,
			Action: ActionRedact, RedactField: "message",
		},
		{
			ID: "default-tag-prod", Name: "Tag Production Sources",
			Enabled: true, MatchSource: "prod",
			Action: ActionTag, TagKey: "env", TagValue: "production",
		},
	}
	for i := range defaults {
		_ = defaults[i].compile()
	}
	e.rules = defaults
	return e
}

// Apply runs the record through all enabled rules.
// Returns (processed record, drop=true if record should be discarded).
func (e *Engine) Apply(rec storage.LogRecord) (storage.LogRecord, bool) {
	e.mu.RLock()
	rules := make([]Rule, len(e.rules))
	copy(rules, e.rules)
	e.mu.RUnlock()

	if rec.Fields == nil {
		rec.Fields = make(map[string]string)
	}

	for _, r := range rules {
		if !r.Enabled || !r.matches(rec) {
			continue
		}
		switch r.Action {
		case ActionDrop:
			return rec, true

		case ActionTag:
			rec.Fields[r.TagKey] = r.TagValue

		case ActionRedact:
			if r.RedactField == "message" {
				if r.compiled != nil {
					rec.Message = r.compiled.ReplaceAllString(rec.Message, "[REDACTED]")
				}
			} else if v, ok := rec.Fields[r.RedactField]; ok && v != "" {
				rec.Fields[r.RedactField] = "[REDACTED]"
			}

		case ActionRoute:
			rec.StreamID = r.RouteStream

		case ActionLevel:
			rec.Level = r.LevelValue

		case ActionParse:
			if r.parseRe != nil {
				m := r.parseRe.FindStringSubmatch(rec.Message)
				names := r.parseRe.SubexpNames()
				for i, name := range names {
					if name != "" && i < len(m) {
						rec.Fields[name] = m[i]
					}
				}
			}
		}
	}
	return rec, false
}

// Rules returns a snapshot of all rules (for the API).
func (e *Engine) Rules() []Rule {
	e.mu.RLock()
	defer e.mu.RUnlock()
	out := make([]Rule, len(e.rules))
	copy(out, e.rules)
	// Zero out compiled fields before serialization.
	for i := range out {
		out[i].compiled = nil
		out[i].parseRe = nil
	}
	return out
}

// SetRules replaces all rules (called from API).
func (e *Engine) SetRules(rules []Rule) error {
	for i := range rules {
		if err := rules[i].compile(); err != nil {
			return err
		}
	}
	e.mu.Lock()
	e.rules = rules
	e.mu.Unlock()
	return nil
}

// AddRule appends a single rule.
func (e *Engine) AddRule(r Rule) error {
	if err := r.compile(); err != nil {
		return err
	}
	e.mu.Lock()
	e.rules = append(e.rules, r)
	e.mu.Unlock()
	return nil
}

// DeleteRule removes a rule by ID.
func (e *Engine) DeleteRule(id string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	filtered := e.rules[:0]
	for _, r := range e.rules {
		if r.ID != id {
			filtered = append(filtered, r)
		}
	}
	e.rules = filtered
}

// MarshalJSON allows the Engine to be serialized directly.
func (e *Engine) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Rules())
}

package pipeline

import (
	"encoding/json"
	"regexp"

	"logdock/internal/storage"
)

// Processor defines the interface for log record transformation.
type Processor interface {
	Process(r *storage.LogRecord) error
}

// Pipeline executes a series of processors in order.
type Pipeline struct {
	Processors []Processor
}

func (p *Pipeline) Execute(r *storage.LogRecord) error {
	for _, proc := range p.Processors {
		if err := proc.Process(r); err != nil {
			return err
		}
	}
	return nil
}

// RegexExtractor extracts fields from the message using a regex with named groups.
type RegexExtractor struct {
	Regex *regexp.Regexp
}

func (e *RegexExtractor) Process(r *storage.LogRecord) error {
	matches := e.Regex.FindStringSubmatch(r.Message)
	if matches == nil {
		return nil
	}
	if r.Fields == nil {
		r.Fields = make(map[string]string)
	}
	for i, name := range e.Regex.SubexpNames() {
		if i != 0 && name != "" {
			r.Fields[name] = matches[i]
		}
	}
	return nil
}

// JSONExtractor attempts to parse the message as JSON and merge it into Fields.
type JSONExtractor struct{}

func (e *JSONExtractor) Process(r *storage.LogRecord) error {
	var data map[string]any
	if err := json.Unmarshal([]byte(r.Message), &data); err != nil {
		return nil // Not JSON, skip
	}
	if r.Fields == nil {
		r.Fields = make(map[string]string)
	}
	for k, v := range data {
		val := jsonValueToString(v)
		r.Fields[k] = val
		// Extract level/source if they appear in the JSON
		if k == "level" && r.Level == "info" {
			r.Level = val
		}
		if k == "source" && r.Source == "unknown" {
			r.Source = val
		}
	}
	return nil
}

func jsonValueToString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	default:
		b, _ := json.Marshal(val)
		return string(b)
	}
}

// StaticField adds a fixed field to the record.
type StaticField struct {
	Key   string
	Value string
}

func (s *StaticField) Process(r *storage.LogRecord) error {
	if r.Fields == nil {
		r.Fields = make(map[string]string)
	}
	r.Fields[s.Key] = s.Value
	return nil
}

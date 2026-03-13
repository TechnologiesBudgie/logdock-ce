// Package settings manages persistent runtime configuration for LogDock.
// Sensitive values (API keys, webhook URLs) are encrypted via CredentialStore.
package settings

import (
	"encoding/json"
	"os"
	"sync"
)

// Settings holds all runtime configuration. Sensitive fields are stored
// separately in the CredentialStore and redacted in JSON output.
type Settings struct {
	mu sync.RWMutex
	s  Snapshot
}

// Snapshot is the serializable (non-secret) view of Settings.
type Snapshot struct {
	InstanceName string            `json:"instance_name"`
	Timezone     string            `json:"timezone"`
	Retention    RetentionConfig   `json:"retention"`
	Storage     StorageConfig     `json:"storage"`
	Security    SecurityConfig    `json:"security"`
	Webhooks    WebhookConfig     `json:"webhooks"`
}

type RetentionConfig struct {
	MaxDays int `json:"maxDays"`
}

type StorageConfig struct {
	DataDir string `json:"dataDir"`
}

type SecurityConfig struct {
	SessionHours int  `json:"sessionHours"`
	RequireMFA   bool `json:"requireMFA"`
	TLSEnabled   bool `json:"tlsEnabled"`
}

type WebhookConfig struct {
	HasSlackURL bool `json:"hasSlackUrl"`
}

// New creates a Settings instance with defaults.
func New() *Settings {
	s := &Settings{}
	s.s = Snapshot{
		InstanceName: "LogDock CE",
		Timezone:     "UTC",
		Retention:   RetentionConfig{MaxDays: 30},
		Storage:     StorageConfig{DataDir: envOr("LOGDOCK_DATA_DIR", "./data")},
		Security:    SecurityConfig{SessionHours: 12},
	}
	return s
}

func (s *Settings) Get() Snapshot {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.s
}

func (s *Settings) Update(snap Snapshot) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.s = snap
}

// MarshalJSON returns a safe (non-secret) JSON snapshot.
func (s *Settings) MarshalJSON() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return json.Marshal(s.s)
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

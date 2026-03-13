package users

import "sync"

// APIKey is a token descriptor suitable for audit views.
type APIKey struct {
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
}

// Profile describes a user account visible to admins.
type Profile struct {
	Username string   `json:"username"`
	Role     string   `json:"role"`
	MFA      bool     `json:"mfa"`
	Keys     []APIKey `json:"keys"`
}

// Service manages user metadata used by admin APIs.
type Service struct {
	mu   sync.Mutex
	data map[string]Profile
}

func New() *Service {
	return &Service{data: map[string]Profile{"admin": {Username: "admin", Role: "admin", MFA: false}}}
}
func (s *Service) List() []Profile {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Profile, 0, len(s.data))
	for _, v := range s.data {
		out = append(out, v)
	}
	return out
}
func (s *Service) Upsert(p Profile) { s.mu.Lock(); defer s.mu.Unlock(); s.data[p.Username] = p }

// Package auth handles JWT issuance/validation, password hashing, token revocation,
// account lockout, and TOTP-based MFA.
// Uses stdlib only — no external dependencies required.
package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User stores identity attributes and security settings.
type User struct {
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"` // bcrypt hash
	Role         string    `json:"role"`
	MFAEnabled   bool      `json:"mfa_enabled"`
	TOTPSecret   string    `json:"-"` // base32 TOTP secret (RFC 6238)
	LockedUntil  time.Time `json:"-"`
	LoginAttempts int      `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	LastLogin    time.Time `json:"last_login"`
}

// loginAttempt tracks bruteforce state per username.
type loginAttempt struct {
	count     int
	lastTry   time.Time
	lockedAt  time.Time
}

// Service handles password verification, JWT issuing, revocation, and role checks.
type Service struct {
	secret      []byte
	mu          sync.RWMutex
	users       map[string]User
	// revokedJTIs is the set of revoked JWT IDs. Pruned hourly.
	revokedJTIs map[string]time.Time
	// loginAttempts tracks per-username failed logins.
	loginAttempts map[string]*loginAttempt

	// Config
	maxLoginAttempts int
	lockoutDuration  time.Duration
	sessionDuration  time.Duration
}

func New(secret string) *Service {
	s := &Service{
		secret:           []byte(secret),
		users:            map[string]User{},
		revokedJTIs:      map[string]time.Time{},
		loginAttempts:    map[string]*loginAttempt{},
		maxLoginAttempts: 10,
		lockoutDuration:  15 * time.Minute,
		sessionDuration:  24 * time.Hour,
	}
	// Prune revoked JTIs every hour to prevent unbounded growth.
	go s.pruneLoop()
	return s
}

func (s *Service) pruneLoop() {
	t := time.NewTicker(time.Hour)
	defer t.Stop()
	for range t.C {
		s.mu.Lock()
		now := time.Now()
		for jti, exp := range s.revokedJTIs {
			if now.After(exp) {
				delete(s.revokedJTIs, jti)
			}
		}
		s.mu.Unlock()
	}
}

func (s *Service) SeedAdmin(password string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users["admin"]; ok {
		return nil
	}
	h, err := hashPassword(password)
	if err != nil {
		return err
	}
	// Generate a TOTP secret for admin (disabled by default).
	totpSecret, _ := generateTOTPSecret()
	s.users["admin"] = User{
		Username:   "admin",
		PasswordHash: h,
		Role:       "admin",
		MFAEnabled: false,
		TOTPSecret: totpSecret,
		CreatedAt:  time.Now().UTC(),
	}
	return nil
}

// Login validates credentials (and MFA when required) and returns a JWT.
// Enforces account lockout after maxLoginAttempts consecutive failures.
func (s *Service) Login(username, password, otp string, globalMFARequired bool) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	u, ok := s.users[username]
	if !ok {
		// Constant-time-ish delay to prevent username enumeration.
		time.Sleep(50 * time.Millisecond)
		return "", errors.New("invalid credentials")
	}

	// Check lockout.
	att := s.loginAttempts[username]
	if att == nil {
		att = &loginAttempt{}
		s.loginAttempts[username] = att
	}
	if !att.lockedAt.IsZero() && time.Since(att.lockedAt) < s.lockoutDuration {
		remaining := s.lockoutDuration - time.Since(att.lockedAt)
		return "", fmt.Errorf("account locked — try again in %d minutes", int(remaining.Minutes())+1)
	}
	// Reset lockout if expired.
	if !att.lockedAt.IsZero() && time.Since(att.lockedAt) >= s.lockoutDuration {
		att.count = 0
		att.lockedAt = time.Time{}
	}

	if err := checkPassword(u.PasswordHash, password); err != nil {
		att.count++
		att.lastTry = time.Now()
		if att.count >= s.maxLoginAttempts {
			att.lockedAt = time.Now()
			return "", fmt.Errorf("too many failed attempts — account locked for %d minutes", int(s.lockoutDuration.Minutes()))
		}
		return "", errors.New("invalid credentials")
	}

	// MFA check.
	if globalMFARequired || u.MFAEnabled {
		if otp == "" {
			return "", errors.New("mfa_required")
		}
		if !s.verifyTOTP(u.TOTPSecret, otp) {
			att.count++
			att.lastTry = time.Now()
			return "", errors.New("invalid MFA code")
		}
	}

	// Successful login — reset attempt counter.
	att.count = 0
	att.lockedAt = time.Time{}

	// Update last login.
	u.LastLogin = time.Now().UTC()
	s.users[username] = u

	return s.issueJWT(username, u.Role)
}

// Logout revokes the given JWT so it cannot be reused.
func (s *Service) Logout(token string) {
	parts := strings.SplitN(token, ".", 3)
	if len(parts) != 3 {
		return
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return
	}
	var claims map[string]any
	if json.Unmarshal(payload, &claims) != nil {
		return
	}
	jti, _ := claims["jti"].(string)
	exp, _ := claims["exp"].(float64)
	if jti == "" {
		return
	}
	s.mu.Lock()
	s.revokedJTIs[jti] = time.Unix(int64(exp), 0)
	s.mu.Unlock()
}

// CreateUser adds a new user. Returns error if username exists.
func (s *Service) CreateUser(username, password, role string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users[username]; ok {
		return errors.New("user already exists")
	}
	h, err := hashPassword(password)
	if err != nil {
		return err
	}
	totpSecret, _ := generateTOTPSecret()
	s.users[username] = User{
		Username:     username,
		PasswordHash: h,
		Role:         role,
		TOTPSecret:   totpSecret,
		CreatedAt:    time.Now().UTC(),
	}
	return nil
}

// DeleteUser removes a user. Cannot delete admin.
func (s *Service) DeleteUser(username string) error {
	if username == "admin" {
		return errors.New("cannot delete admin")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.users, username)
	return nil
}

// ChangePassword updates a user's password and revokes all existing sessions.
func (s *Service) ChangePassword(username, newPassword string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, ok := s.users[username]
	if !ok {
		return errors.New("user not found")
	}
	h, err := hashPassword(newPassword)
	if err != nil {
		return err
	}
	u.PasswordHash = h
	s.users[username] = u
	return nil
}

// EnableMFA enables TOTP MFA for a user and returns the provisioning URI.
func (s *Service) EnableMFA(username string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, ok := s.users[username]
	if !ok {
		return "", errors.New("user not found")
	}
	secret, err := generateTOTPSecret()
	if err != nil {
		return "", err
	}
	u.TOTPSecret = secret
	u.MFAEnabled = true
	s.users[username] = u
	// Return otpauth URI for QR code generation.
	uri := fmt.Sprintf("otpauth://totp/LogDock:%s?secret=%s&issuer=LogDock&algorithm=SHA1&digits=6&period=30",
		username, secret)
	return uri, nil
}

// DisableMFA disables TOTP for a user.
func (s *Service) DisableMFA(username string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, ok := s.users[username]
	if !ok {
		return errors.New("user not found")
	}
	u.MFAEnabled = false
	s.users[username] = u
	return nil
}

// UserList returns all users (safe view: no password hashes).
func (s *Service) UserList() []map[string]any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]map[string]any, 0, len(s.users))
	for _, u := range s.users {
		att := s.loginAttempts[u.Username]
		locked := false
		if att != nil && !att.lockedAt.IsZero() && time.Since(att.lockedAt) < s.lockoutDuration {
			locked = true
		}
		out = append(out, map[string]any{
			"username":    u.Username,
			"role":        u.Role,
			"mfa_enabled": u.MFAEnabled,
			"created_at":  u.CreatedAt,
			"last_login":  u.LastLogin,
			"locked":      locked,
		})
	}
	return out
}

// UnlockUser clears the lockout for a user.
func (s *Service) UnlockUser(username string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.loginAttempts, username)
}

// SetMaxLoginAttempts configures lockout threshold.
func (s *Service) SetMaxLoginAttempts(n int) {
	s.mu.Lock()
	s.maxLoginAttempts = n
	s.mu.Unlock()
}

// Authorize checks token validity and role membership.
func (s *Service) Authorize(token string, roles ...string) error {
	claims, err := s.parseJWT(token)
	if err != nil {
		return err
	}
	if len(roles) == 0 {
		return nil
	}
	role, _ := claims["role"].(string)
	for _, r := range roles {
		if role == r {
			return nil
		}
	}
	return errors.New("forbidden")
}

// issueJWT creates a signed HS256 JWT with a unique JTI for revocation support.
func (s *Service) issueJWT(username, role string) (string, error) {
	jtiBytes := make([]byte, 16)
	rand.Read(jtiBytes)
	jti := hex.EncodeToString(jtiBytes)

	type jwtClaims struct {
		Sub  string `json:"sub"`
		Role string `json:"role"`
		JTI  string `json:"jti"`
		Exp  int64  `json:"exp"`
		Iat  int64  `json:"iat"`
	}
	header, _ := json.Marshal(map[string]string{"alg": "HS256", "typ": "JWT"})
	payload, _ := json.Marshal(jwtClaims{
		Sub: username, Role: role, JTI: jti,
		Exp: time.Now().Add(s.sessionDuration).Unix(),
		Iat: time.Now().Unix(),
	})
	h := base64.RawURLEncoding.EncodeToString(header)
	p := base64.RawURLEncoding.EncodeToString(payload)
	sig := jwtSign(s.secret, h+"."+p)
	return h + "." + p + "." + sig, nil
}

// parseJWT validates and decodes a JWT. Checks revocation list.
func (s *Service) parseJWT(token string) (map[string]any, error) {
	parts := strings.SplitN(token, ".", 3)
	if len(parts) != 3 {
		return nil, errors.New("invalid token")
	}
	expected := jwtSign(s.secret, parts[0]+"."+parts[1])
	if !hmac.Equal([]byte(expected), []byte(parts[2])) {
		return nil, errors.New("invalid token")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, errors.New("invalid token")
	}
	var claims map[string]any
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, errors.New("invalid token")
	}
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, errors.New("token expired")
		}
	}
	// Revocation check.
	if jti, ok := claims["jti"].(string); ok && jti != "" {
		s.mu.RLock()
		_, revoked := s.revokedJTIs[jti]
		s.mu.RUnlock()
		if revoked {
			return nil, errors.New("token revoked")
		}
	}
	return claims, nil
}

func jwtSign(secret []byte, msg string) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(msg))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

// hashPassword creates a secure bcrypt hash.
func hashPassword(plain string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashPassword: %w", err)
	}
	return string(h), nil
}

// checkPassword verifies a plain password against a stored bcrypt hash.
func checkPassword(storedHash, plain string) error {
	if strings.Contains(storedHash, ":") {
		// Migration fallback for old HMAC hashes
		parts := strings.SplitN(storedHash, ":", 2)
		if len(parts) == 2 {
			salt, _ := hex.DecodeString(parts[0])
			expected, _ := hex.DecodeString(parts[1])
			mac := hmac.New(sha256.New, salt)
			mac.Write([]byte(plain))
			if hmac.Equal(mac.Sum(nil), expected) {
				return nil
			}
		}
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(plain))
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}

// generateTOTPSecret returns a random base32-encoded TOTP secret.
func generateTOTPSecret() (string, error) {
	b := make([]byte, 20)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base32.StdEncoding.EncodeToString(b), nil
}

// verifyTOTP checks a 6-digit TOTP code against a base32 secret (RFC 6238, SHA1, 30s window).
// Accepts current and adjacent time steps (±1) to account for clock drift.
func (s *Service) verifyTOTP(secret, code string) bool {
	key, err := base32.StdEncoding.DecodeString(strings.ToUpper(strings.TrimRight(secret, "=")))
	if err != nil {
		return false
	}
	now := time.Now().Unix() / 30
	for _, step := range []int64{-1, 0, 1} {
		if totpCode(key, now+step) == code {
			return true
		}
	}
	return false
}

func totpCode(key []byte, counter int64) string {
	msg := make([]byte, 8)
	binary.BigEndian.PutUint64(msg, uint64(counter))
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	h := mac.Sum(nil)
	offset := h[len(h)-1] & 0x0f
	code := (binary.BigEndian.Uint32(h[offset:offset+4]) & 0x7fffffff) % 1000000
	return fmt.Sprintf("%06d", code)
}

// Parse decodes a JWT without role checking (for reading claims).
func (s *Service) Parse(token string) (map[string]any, error) {
	return s.parseJWT(token)
}

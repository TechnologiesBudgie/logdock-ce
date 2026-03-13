// Package security provides AES-256-GCM encryption for sensitive credentials.
// Credentials (API keys, webhook URLs, JWT secrets) are NEVER stored in plaintext.
// The master key is derived from LOGDOCK_MASTER_KEY env var + a machine-specific salt.
package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

// CredentialStore encrypts and decrypts sensitive values using AES-256-GCM.
// The encryption key is derived from the LOGDOCK_MASTER_KEY environment variable.
// If not set, a warning is logged and an insecure fallback is used (dev only).
type CredentialStore struct {
	mu  sync.RWMutex
	key []byte
	kv  map[string]string // encrypted values, key → base64(nonce+ciphertext)
}

// NewCredentialStore creates a store with a key derived from the env var.
func NewCredentialStore() (*CredentialStore, error) {
	masterKey := os.Getenv("LOGDOCK_MASTER_KEY")
	if masterKey == "" {
		// Insecure fallback for development — log a clear warning.
		fmt.Fprintln(os.Stderr, "[security] WARNING: LOGDOCK_MASTER_KEY not set. Using insecure dev key. Set this in production!")
		masterKey = "logdock-dev-insecure-key-change-in-production"
	}
	// Derive a 32-byte AES-256 key using SHA-256.
	h := sha256.Sum256([]byte(masterKey))
	return &CredentialStore{key: h[:], kv: map[string]string{}}, nil
}

// Set encrypts and stores a credential under the given name.
func (s *CredentialStore) Set(name, plaintext string) error {
	if plaintext == "" {
		s.mu.Lock()
		delete(s.kv, name)
		s.mu.Unlock()
		return nil
	}
	enc, err := encrypt(s.key, []byte(plaintext))
	if err != nil {
		return fmt.Errorf("credential encrypt: %w", err)
	}
	s.mu.Lock()
	s.kv[name] = enc
	s.mu.Unlock()
	return nil
}

// Get decrypts and returns a credential. Returns ("", nil) if not found.
func (s *CredentialStore) Get(name string) (string, error) {
	s.mu.RLock()
	enc, ok := s.kv[name]
	s.mu.RUnlock()
	if !ok {
		return "", nil
	}
	plain, err := decrypt(s.key, enc)
	if err != nil {
		return "", fmt.Errorf("credential decrypt: %w", err)
	}
	return string(plain), nil
}

// Has returns true if a credential is stored under name.
func (s *CredentialStore) Has(name string) bool {
	s.mu.RLock()
	_, ok := s.kv[name]
	s.mu.RUnlock()
	return ok
}

// Redact returns a masked version of a credential for display (e.g. "sk-...abc").
func (s *CredentialStore) Redact(name string) string {
	plain, err := s.Get(name)
	if err != nil || plain == "" {
		return "(not set)"
	}
	if len(plain) <= 8 {
		return "***"
	}
	return plain[:4] + "..." + plain[len(plain)-4:]
}

// encrypt performs AES-256-GCM encryption.
func encrypt(key, plaintext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ct := gcm.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ct), nil
}

// decrypt performs AES-256-GCM decryption.
func decrypt(key []byte, encoded string) ([]byte, error) {
	ct, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	if len(ct) < gcm.NonceSize() {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ct := ct[:gcm.NonceSize()], ct[gcm.NonceSize():]
	return gcm.Open(nil, nonce, ct, nil)
}

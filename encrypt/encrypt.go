// Package encrypt provides a DI-friendly AES-256-GCM string encryption service.
package encrypt

import (
	"errors"

	"github.com/andywang2011/go-libs/cryptox/aesx"
)

// KeySize is the required key length in bytes (AES-256).
const KeySize = 32

// ErrInvalidKeySize is returned when EncryptConfig.Key is not exactly KeySize bytes.
var ErrInvalidKeySize = errors.New("encrypt: key must be 32 bytes (AES-256)")

// EncryptConfig holds AES-256-GCM encryption settings.
type EncryptConfig struct {
	// Key is the raw AES-256 key; it must be exactly 32 bytes.
	Key string
}

// EncryptService performs AES-256-GCM encrypt/decrypt with an injected key.
// Ciphertext strings are base64(nonce || ciphertext || tag).
type EncryptService struct {
	Key string
}

// NewEncryptService copies the key from an existing service.
func NewEncryptService(e *EncryptService) *EncryptService {
	return &EncryptService{Key: e.Key}
}

// NewEncryptServiceFromConfig constructs a service from config (Wire / tests).
// It returns ErrInvalidKeySize when the key is not 32 bytes.
func NewEncryptServiceFromConfig(cfg EncryptConfig) (*EncryptService, error) {
	if len(cfg.Key) != KeySize {
		return nil, ErrInvalidKeySize
	}
	return &EncryptService{Key: cfg.Key}, nil
}

func (e *EncryptService) key() ([]byte, error) {
	if len(e.Key) != KeySize {
		return nil, ErrInvalidKeySize
	}
	return []byte(e.Key), nil
}

// Encrypt returns base64(nonce || ciphertext || tag); empty input yields empty output.
func (e *EncryptService) Encrypt(plainText string) (string, error) {
	if len(plainText) == 0 {
		return "", nil
	}
	key, err := e.key()
	if err != nil {
		return "", err
	}
	return aesx.GCMEncryptString(key, plainText)
}

// Decrypt reverses Encrypt and verifies authenticity; empty input yields empty output.
func (e *EncryptService) Decrypt(cipherText string) (string, error) {
	if len(cipherText) == 0 {
		return "", nil
	}
	key, err := e.key()
	if err != nil {
		return "", err
	}
	return aesx.GCMDecryptString(key, cipherText)
}

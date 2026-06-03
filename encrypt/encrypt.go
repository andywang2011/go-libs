package encrypt

import (
	"github.com/andywang2011/go-libs/cryptox/aesx"
)

// EncryptConfig holds AES-CBC encryption settings.
type EncryptConfig struct {
	Iv  string
	Key string
}

// EncryptService performs AES-CBC encrypt/decrypt with injected IV and key.
type EncryptService struct {
	Iv  string
	Key string
}

// NewEncryptService copies fields from an existing service (legacy compatibility).
func NewEncryptService(e *EncryptService) *EncryptService {
	return &EncryptService{Iv: e.Iv, Key: e.Key}
}

// NewEncryptServiceFromConfig constructs a service from config (Wire / tests).
func NewEncryptServiceFromConfig(cfg EncryptConfig) *EncryptService {
	return &EncryptService{Iv: cfg.Iv, Key: cfg.Key}
}

func (e *EncryptService) Encrypt(plainText string) (string, error) {
	if len(plainText) == 0 {
		return "", nil
	}
	return aesx.CBCEncryptWithIV(e.Iv, e.Key, plainText)
}

func (e *EncryptService) Decrypt(cipherText string) (string, error) {
	if len(cipherText) == 0 {
		return "", nil
	}
	return aesx.CBCDecryptWithIV(e.Iv, e.Key, cipherText)
}

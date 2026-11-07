package aesx

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

// ErrCiphertextTooShort is returned when GCM input is shorter than nonce plus tag.
var ErrCiphertextTooShort = errors.New("aesx: ciphertext too short")

func newGCM(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key) // key must be 16, 24, or 32 bytes
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

// GCMEncrypt encrypts plaintext with AES-GCM using a random nonce.
// The key must be 16, 24, or 32 bytes. Output layout: nonce || ciphertext || tag.
func GCMEncrypt(key, plaintext []byte) ([]byte, error) {
	gcm, err := newGCM(key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize(), gcm.NonceSize()+len(plaintext)+gcm.Overhead())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// GCMDecrypt decrypts data produced by GCMEncrypt and verifies the authentication tag.
func GCMDecrypt(key, ciphertext []byte) ([]byte, error) {
	gcm, err := newGCM(key)
	if err != nil {
		return nil, err
	}
	ns := gcm.NonceSize()
	if len(ciphertext) < ns+gcm.Overhead() {
		return nil, ErrCiphertextTooShort
	}
	nonce, body := ciphertext[:ns], ciphertext[ns:]
	return gcm.Open(nil, nonce, body, nil)
}

// GCMEncryptString encrypts plainText and returns base64(nonce || ciphertext || tag).
func GCMEncryptString(key []byte, plainText string) (string, error) {
	out, err := GCMEncrypt(key, []byte(plainText))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(out), nil
}

// GCMDecryptString decrypts base64(nonce || ciphertext || tag) produced by GCMEncryptString.
func GCMDecryptString(key []byte, cipherText string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	plain, err := GCMDecrypt(key, raw)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}

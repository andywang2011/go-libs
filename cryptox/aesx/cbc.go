// Package aesx provides AES helpers: AES-GCM (preferred) and AES-CBC with an explicit IV.
//
// ECB mode and CBC with the key reused as the IV are intentionally not provided.
// CBC offers no integrity protection; prefer GCMEncrypt / GCMDecrypt for new code.
package aesx

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// ErrInvalidIV is returned when the CBC IV length is not the AES block size.
var ErrInvalidIV = errors.New("aesx: IV must be 16 bytes")

// ErrInvalidCiphertext is returned when CBC ciphertext length is not a positive multiple of the block size.
var ErrInvalidCiphertext = errors.New("aesx: invalid ciphertext length")

// CBCDecryptWithIV decrypts base64 ciphertext using base64-encoded iv and key.
// It returns an error if the padding is invalid.
func CBCDecryptWithIV(iv, key, cipherText string) (string, error) {
	ivByt, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	keyByt, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	encryptedByt, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(keyByt)
	if err != nil {
		return "", err
	}
	if len(ivByt) != block.BlockSize() {
		return "", ErrInvalidIV
	}
	if len(encryptedByt) == 0 || len(encryptedByt)%block.BlockSize() != 0 {
		return "", ErrInvalidCiphertext
	}
	plainByt := make([]byte, len(encryptedByt))
	cipher.NewCBCDecrypter(block, ivByt).CryptBlocks(plainByt, encryptedByt)

	b, err := PKCS7UnPadding(plainByt)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// CBCEncryptWithIV encrypts plainText using base64-encoded iv and key and returns base64 ciphertext.
func CBCEncryptWithIV(iv, key, plainText string) (string, error) {
	ivByt, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	keyByt, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyByt)
	if err != nil {
		return "", err
	}
	if len(ivByt) != block.BlockSize() {
		return "", ErrInvalidIV
	}
	plainByt := PKCS7Padding([]byte(plainText), block.BlockSize())
	crypted := make([]byte, len(plainByt))
	cipher.NewCBCEncrypter(block, ivByt).CryptBlocks(crypted, plainByt)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

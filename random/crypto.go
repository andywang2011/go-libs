package random

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

// Bytes returns n cryptographically random bytes.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func UUID() string {
	b, err := Bytes(16)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func UUIDWithLen(n int) string {
	b, err := Bytes(n / 2)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

// NumCode returns a cryptographically random six-digit numeric string.
func NumCode() string {
	num, _ := rand.Int(rand.Reader, big.NewInt(1e6))
	return fmt.Sprintf("%06d", num.Int64())
}

// Package random provides cryptographically secure random helpers.
package random

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
)

// ErrInvalidLength is returned when a requested length is negative.
var ErrInvalidLength = errors.New("random: length must be non-negative")

// randRead is a seam for tests to simulate CSPRNG failure.
var randRead = rand.Read

// Bytes returns n cryptographically random bytes.
func Bytes(n int) ([]byte, error) {
	if n < 0 {
		return nil, ErrInvalidLength
	}
	b := make([]byte, n)
	if _, err := randRead(b); err != nil {
		return nil, err
	}
	return b, nil
}

// UUID returns 32 lowercase hex characters (128 random bits).
func UUID() (string, error) {
	return UUIDWithLen(32)
}

// UUIDWithLen returns exactly n lowercase hex characters from the CSPRNG.
func UUIDWithLen(n int) (string, error) {
	if n < 0 {
		return "", ErrInvalidLength
	}
	b, err := Bytes((n + 1) / 2)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b)[:n], nil
}

// MustToken returns n random hex characters and panics if the CSPRNG fails
// or n is negative. Intended for request IDs where failure is unrecoverable.
func MustToken(n int) string {
	s, err := UUIDWithLen(n)
	if err != nil {
		panic(fmt.Sprintf("random: MustToken: %v", err))
	}
	return s
}

// NumCode returns a cryptographically random six-digit numeric string.
func NumCode() (string, error) {
	num, err := rand.Int(rand.Reader, big.NewInt(1e6))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", num.Int64()), nil
}

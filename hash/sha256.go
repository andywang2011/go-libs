// Package hash provides SHA-256 and HMAC-SHA256 helpers.
package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// SHA256Hex returns the lowercase hex SHA-256 digest of str.
func SHA256Hex(str string) string {
	sum := sha256.Sum256([]byte(str))
	return hex.EncodeToString(sum[:])
}

// HMACSHA256 returns the HMAC-SHA256 of msg under key.
func HMACSHA256(key, msg []byte) []byte {
	m := hmac.New(sha256.New, key)
	m.Write(msg)
	return m.Sum(nil)
}

// HMACSHA256Hex returns the lowercase hex HMAC-SHA256 of msg under key.
func HMACSHA256Hex(key, msg []byte) string {
	return hex.EncodeToString(HMACSHA256(key, msg))
}

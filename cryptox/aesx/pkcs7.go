package aesx

import (
	"bytes"
	"crypto/aes"
	"errors"
)

// ErrInvalidPadding is returned when PKCS#7 padding is malformed.
var ErrInvalidPadding = errors.New("aesx: invalid PKCS7 padding")

// PKCS7Padding applies PKCS#7 padding to data for the given block size.
// The input slice is not modified.
func PKCS7Padding(data []byte, blockSize int) []byte {
	if blockSize <= 0 || blockSize > 255 {
		panic("aesx: PKCS7 blockSize must be in 1..255")
	}
	padding := blockSize - len(data)%blockSize
	padByte := byte(padding)
	padtext := bytes.Repeat([]byte{padByte}, padding)
	out := make([]byte, 0, len(data)+padding)
	out = append(out, data...)
	return append(out, padtext...)
}

// PKCS7UnPadding removes PKCS#7 padding (AES block size) and returns an error
// when the padding is empty, out of range, or inconsistent.
func PKCS7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 || length%aes.BlockSize != 0 {
		return nil, ErrInvalidPadding
	}
	pad := int(data[length-1])
	if pad == 0 || pad > aes.BlockSize || pad > length {
		return nil, ErrInvalidPadding
	}
	for _, b := range data[length-pad:] {
		if int(b) != pad {
			return nil, ErrInvalidPadding
		}
	}
	return data[:length-pad], nil
}

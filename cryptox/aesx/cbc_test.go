package aesx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"testing"
)

var (
	cbcKeyB64 = base64.StdEncoding.EncodeToString([]byte("0123456789abcdef"))
	cbcIVB64  = base64.StdEncoding.EncodeToString([]byte("fedcba9876543210"))
)

func TestCBCWithIVRoundTrip(t *testing.T) {
	for _, plain := range []string{"", "a", "verystart", "exactly16bytes!!", "longer than one block of plaintext"} {
		enc, err := CBCEncryptWithIV(cbcIVB64, cbcKeyB64, plain)
		if err != nil {
			t.Fatalf("encrypt %q: %v", plain, err)
		}
		dec, err := CBCDecryptWithIV(cbcIVB64, cbcKeyB64, enc)
		if err != nil {
			t.Fatalf("decrypt %q: %v", plain, err)
		}
		if dec != plain {
			t.Fatalf("got %q want %q", dec, plain)
		}
	}
}

func TestCBCWithIVErrors(t *testing.T) {
	good, err := CBCEncryptWithIV(cbcIVB64, cbcKeyB64, "hello")
	if err != nil {
		t.Fatal(err)
	}
	badKey := base64.StdEncoding.EncodeToString([]byte("short"))
	shortIV := base64.StdEncoding.EncodeToString([]byte("iv"))

	if _, err := CBCEncryptWithIV("!!", cbcKeyB64, "x"); err == nil {
		t.Error("encrypt: expected bad iv base64")
	}
	if _, err := CBCEncryptWithIV(cbcIVB64, "!!", "x"); err == nil {
		t.Error("encrypt: expected bad key base64")
	}
	if _, err := CBCEncryptWithIV(cbcIVB64, badKey, "x"); err == nil {
		t.Error("encrypt: expected bad key size")
	}
	if _, err := CBCEncryptWithIV(shortIV, cbcKeyB64, "x"); !errors.Is(err, ErrInvalidIV) {
		t.Errorf("encrypt: expected ErrInvalidIV, got %v", err)
	}

	if _, err := CBCDecryptWithIV("!!", cbcKeyB64, good); err == nil {
		t.Error("decrypt: expected bad iv base64")
	}
	if _, err := CBCDecryptWithIV(cbcIVB64, "!!", good); err == nil {
		t.Error("decrypt: expected bad key base64")
	}
	if _, err := CBCDecryptWithIV(cbcIVB64, cbcKeyB64, "!!"); err == nil {
		t.Error("decrypt: expected bad ciphertext base64")
	}
	if _, err := CBCDecryptWithIV(cbcIVB64, badKey, good); err == nil {
		t.Error("decrypt: expected bad key size")
	}
	if _, err := CBCDecryptWithIV(shortIV, cbcKeyB64, good); !errors.Is(err, ErrInvalidIV) {
		t.Errorf("decrypt: expected ErrInvalidIV, got %v", err)
	}
	if _, err := CBCDecryptWithIV(cbcIVB64, cbcKeyB64, ""); !errors.Is(err, ErrInvalidCiphertext) {
		t.Errorf("decrypt empty: got %v", err)
	}
	odd := base64.StdEncoding.EncodeToString([]byte("not-block-aligned"))
	if _, err := CBCDecryptWithIV(cbcIVB64, cbcKeyB64, odd); !errors.Is(err, ErrInvalidCiphertext) {
		t.Errorf("decrypt unaligned: got %v", err)
	}
}

func TestCBCDecryptBadPadding(t *testing.T) {
	keyByt, _ := base64.StdEncoding.DecodeString(cbcKeyB64)
	ivByt, _ := base64.StdEncoding.DecodeString(cbcIVB64)
	block, err := aes.NewCipher(keyByt)
	if err != nil {
		t.Fatal(err)
	}
	for name, plain := range map[string][]byte{
		"zero pad":     append(bytes.Repeat([]byte{'a'}, 15), 0),
		"pad too big":  bytes.Repeat([]byte{17}, 16),
		"inconsistent": append(bytes.Repeat([]byte{'a'}, 14), 1, 2),
	} {
		ct := make([]byte, len(plain))
		cipher.NewCBCEncrypter(block, ivByt).CryptBlocks(ct, plain)
		_, err := CBCDecryptWithIV(cbcIVB64, cbcKeyB64, base64.StdEncoding.EncodeToString(ct))
		if !errors.Is(err, ErrInvalidPadding) {
			t.Errorf("%s: expected ErrInvalidPadding, got %v", name, err)
		}
	}
}

func TestPKCS7Padding(t *testing.T) {
	in := []byte("abc")
	out := PKCS7Padding(in, 16)
	if len(out) != 16 || out[15] != 13 {
		t.Fatalf("unexpected padding: %v", out)
	}
	if string(in) != "abc" {
		t.Fatal("input mutated")
	}
	full := PKCS7Padding(make([]byte, 16), 16)
	if len(full) != 32 || full[31] != 16 {
		t.Fatalf("full block padding: %v", full)
	}
}

func TestPKCS7UnPadding(t *testing.T) {
	got, err := PKCS7UnPadding(PKCS7Padding([]byte("abc"), 16))
	if err != nil || string(got) != "abc" {
		t.Fatalf("got %q err %v", got, err)
	}
	got, err = PKCS7UnPadding(PKCS7Padding(nil, 16))
	if err != nil || len(got) != 0 {
		t.Fatalf("empty: got %q err %v", got, err)
	}

	bad := map[string][]byte{
		"nil":           nil,
		"unaligned":     make([]byte, 15),
		"zero pad":      append(make([]byte, 15), 0),
		"pad too large": append(make([]byte, 15), 17),
		"inconsistent":  append(append(make([]byte, 13), 1, 3), 3),
		"pad > length":  append(make([]byte, 15), 16)[:16],
	}
	// "pad > length" above equals a 16-byte block with pad 16 but non-matching bytes.
	bad["pad > length"][0] = 1
	for name, in := range bad {
		if _, err := PKCS7UnPadding(in); !errors.Is(err, ErrInvalidPadding) {
			t.Errorf("%s: expected ErrInvalidPadding, got %v", name, err)
		}
	}
}

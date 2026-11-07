package aesx

import (
	"bytes"
	"crypto/rand"
	"errors"
	"testing"
)

func testKey(t *testing.T, n int) []byte {
	t.Helper()
	k := make([]byte, n)
	if _, err := rand.Read(k); err != nil {
		t.Fatal(err)
	}
	return k
}

func TestGCMRoundTripKeySizes(t *testing.T) {
	for _, size := range []int{16, 24, 32} {
		key := testKey(t, size)
		for _, plain := range [][]byte{nil, {}, []byte("a"), []byte("hello gcm"), bytes.Repeat([]byte("x"), 4096)} {
			enc, err := GCMEncrypt(key, plain)
			if err != nil {
				t.Fatalf("size %d encrypt: %v", size, err)
			}
			if len(enc) != 12+len(plain)+16 {
				t.Fatalf("size %d: unexpected length %d", size, len(enc))
			}
			dec, err := GCMDecrypt(key, enc)
			if err != nil {
				t.Fatalf("size %d decrypt: %v", size, err)
			}
			if !bytes.Equal(dec, plain) {
				t.Fatalf("size %d: roundtrip mismatch", size)
			}
		}
	}
}

func TestGCMRandomNonce(t *testing.T) {
	key := testKey(t, 32)
	a, _ := GCMEncrypt(key, []byte("same"))
	b, _ := GCMEncrypt(key, []byte("same"))
	if bytes.Equal(a, b) {
		t.Fatal("expected distinct ciphertexts for same plaintext")
	}
	if bytes.Equal(a[:12], b[:12]) {
		t.Fatal("expected distinct nonces")
	}
}

func TestGCMTamperDetection(t *testing.T) {
	key := testKey(t, 32)
	enc, err := GCMEncrypt(key, []byte("sensitive"))
	if err != nil {
		t.Fatal(err)
	}
	// Flip one bit in nonce, body, and tag regions.
	for _, idx := range []int{0, 12, len(enc) - 1} {
		tampered := append([]byte(nil), enc...)
		tampered[idx] ^= 0x01
		if _, err := GCMDecrypt(key, tampered); err == nil {
			t.Fatalf("expected tamper error at byte %d", idx)
		}
	}
	// Truncation.
	if _, err := GCMDecrypt(key, enc[:len(enc)-1]); err == nil {
		t.Fatal("expected error on truncated ciphertext")
	}
}

func TestGCMWrongKey(t *testing.T) {
	enc, _ := GCMEncrypt(testKey(t, 32), []byte("x"))
	if _, err := GCMDecrypt(testKey(t, 32), enc); err == nil {
		t.Fatal("expected error with wrong key")
	}
}

func TestGCMBadKeySize(t *testing.T) {
	for _, n := range []int{0, 1, 15, 17, 31, 33, 64} {
		key := make([]byte, n)
		if _, err := GCMEncrypt(key, []byte("x")); err == nil {
			t.Fatalf("encrypt: expected error for key size %d", n)
		}
		if _, err := GCMDecrypt(key, make([]byte, 64)); err == nil {
			t.Fatalf("decrypt: expected error for key size %d", n)
		}
	}
}

func TestGCMDecryptTooShort(t *testing.T) {
	key := testKey(t, 32)
	for _, n := range []int{0, 1, 12, 27} {
		_, err := GCMDecrypt(key, make([]byte, n))
		if !errors.Is(err, ErrCiphertextTooShort) {
			t.Fatalf("len %d: got %v", n, err)
		}
	}
}

func TestGCMStringRoundTrip(t *testing.T) {
	key := testKey(t, 32)
	enc, err := GCMEncryptString(key, "héllo 世界")
	if err != nil {
		t.Fatal(err)
	}
	dec, err := GCMDecryptString(key, enc)
	if err != nil {
		t.Fatal(err)
	}
	if dec != "héllo 世界" {
		t.Fatalf("got %q", dec)
	}
}

func TestGCMStringErrors(t *testing.T) {
	key := testKey(t, 32)
	if _, err := GCMEncryptString([]byte("short"), "x"); err == nil {
		t.Fatal("expected bad key error")
	}
	if _, err := GCMDecryptString(key, "!!!not base64!!!"); err == nil {
		t.Fatal("expected base64 error")
	}
	if _, err := GCMDecryptString(key, "AAAA"); err == nil {
		t.Fatal("expected too short error")
	}
}

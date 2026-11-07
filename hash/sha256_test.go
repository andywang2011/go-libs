package hash

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestSHA256Hex(t *testing.T) {
	cases := map[string]string{
		"":      "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		"abc":   "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
		"hello": "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
	}
	for in, want := range cases {
		if got := SHA256Hex(in); got != want {
			t.Errorf("SHA256Hex(%q) = %s, want %s", in, got, want)
		}
	}
}

func TestHMACSHA256RFC4231(t *testing.T) {
	// RFC 4231 test case 2.
	key := []byte("Jefe")
	msg := []byte("what do ya want for nothing?")
	want := "5bdcc146bf60754e6a042426089575c75a003f089d2739839dec58b964ec3843"
	got := HMACSHA256(key, msg)
	if hex.EncodeToString(got) != want {
		t.Fatalf("HMACSHA256 = %x", got)
	}
	if HMACSHA256Hex(key, msg) != want {
		t.Fatal("HMACSHA256Hex mismatch")
	}
}

func TestHMACSHA256KeySensitivity(t *testing.T) {
	msg := []byte("m")
	if bytes.Equal(HMACSHA256([]byte("k1"), msg), HMACSHA256([]byte("k2"), msg)) {
		t.Fatal("different keys must differ")
	}
	if len(HMACSHA256(nil, nil)) != 32 {
		t.Fatal("expected 32-byte MAC")
	}
}

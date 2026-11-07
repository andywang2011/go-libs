package random_test

import (
	"encoding/hex"
	"regexp"
	"testing"
	"unicode"

	"github.com/andywang2011/go-libs/random"
)

func TestBytesLength(t *testing.T) {
	for _, n := range []int{0, 1, 16, 32} {
		b, err := random.Bytes(n)
		if err != nil {
			t.Fatalf("Bytes(%d): %v", n, err)
		}
		if len(b) != n {
			t.Fatalf("Bytes(%d): len %d", n, len(b))
		}
	}
}

func TestUUIDFormats(t *testing.T) {
	id := random.UUID()
	if id == "" {
		t.Fatal("UUID returned empty")
	}
	if len(id) != 32 {
		t.Fatalf("UUID len: %d", len(id))
	}
	if _, err := hex.DecodeString(id); err != nil {
		t.Fatalf("UUID hex: %v", err)
	}

	short := random.UUIDWithLen(8)
	if short == "" {
		t.Fatal("UUIDWithLen returned empty")
	}
	if len(short) != 8 {
		t.Fatalf("UUIDWithLen(8): %q", short)
	}
	if _, err := hex.DecodeString(short); err != nil {
		t.Fatalf("UUIDWithLen hex: %v", err)
	}
}

func TestNumCodeFormat(t *testing.T) {
	re := regexp.MustCompile(`^\d{6}$`)
	for i := 0; i < 20; i++ {
		code := random.NumCode()
		if !re.MatchString(code) {
			t.Fatalf("NumCode: %q", code)
		}
	}
}

func TestUUIDDistinct(t *testing.T) {
	a := random.UUID()
	b := random.UUID()
	if a == b {
		t.Fatal("expected distinct UUIDs")
	}
}

func TestNumCodeNumericOnly(t *testing.T) {
	for _, r := range random.NumCode() {
		if !unicode.IsDigit(r) {
			t.Fatalf("non-digit in NumCode")
		}
	}
}

package random

import (
	"encoding/hex"
	"errors"
	"regexp"
	"sort"
	"testing"
	"unicode"
)

func TestBytesLength(t *testing.T) {
	for _, n := range []int{0, 1, 16, 32} {
		b, err := Bytes(n)
		if err != nil {
			t.Fatalf("Bytes(%d): %v", n, err)
		}
		if len(b) != n {
			t.Fatalf("Bytes(%d): len %d", n, len(b))
		}
	}
}

func TestBytesNegative(t *testing.T) {
	if _, err := Bytes(-1); !errors.Is(err, ErrInvalidLength) {
		t.Fatalf("got %v", err)
	}
}

func TestUUIDFormats(t *testing.T) {
	id, err := UUID()
	if err != nil {
		t.Fatal(err)
	}
	if len(id) != 32 {
		t.Fatalf("UUID len: %d", len(id))
	}
	if _, err := hex.DecodeString(id); err != nil {
		t.Fatalf("UUID hex: %v", err)
	}
}

func TestUUIDWithLen(t *testing.T) {
	for _, n := range []int{0, 1, 8, 10, 32, 43} {
		s, err := UUIDWithLen(n)
		if err != nil {
			t.Fatalf("UUIDWithLen(%d): %v", n, err)
		}
		if len(s) != n {
			t.Fatalf("UUIDWithLen(%d): got len %d", n, len(s))
		}
		if _, err := hex.DecodeString(s + pad(n)); err != nil {
			t.Fatalf("UUIDWithLen(%d) hex: %v", n, err)
		}
	}
	if _, err := UUIDWithLen(-1); !errors.Is(err, ErrInvalidLength) {
		t.Fatalf("negative: %v", err)
	}
}

func pad(n int) string {
	if n%2 == 1 {
		return "0"
	}
	return ""
}

func TestUUIDDistinct(t *testing.T) {
	a, _ := UUID()
	b, _ := UUID()
	if a == b {
		t.Fatal("expected distinct UUIDs")
	}
}

func TestNumCodeFormat(t *testing.T) {
	re := regexp.MustCompile(`^\d{6}$`)
	for i := 0; i < 20; i++ {
		code, err := NumCode()
		if err != nil {
			t.Fatal(err)
		}
		if !re.MatchString(code) {
			t.Fatalf("NumCode: %q", code)
		}
		for _, r := range code {
			if !unicode.IsDigit(r) {
				t.Fatalf("non-digit in NumCode")
			}
		}
	}
}

func TestMustToken(t *testing.T) {
	tok := MustToken(24)
	if len(tok) != 24 {
		t.Fatalf("len %d", len(tok))
	}
	if MustToken(24) == tok {
		t.Fatal("expected distinct tokens")
	}
}

func TestMustTokenPanicsOnNegative(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	MustToken(-1)
}

func TestFailuresPropagate(t *testing.T) {
	orig := randRead
	defer func() { randRead = orig }()
	boom := errors.New("csprng down")
	randRead = func([]byte) (int, error) { return 0, boom }

	if _, err := Bytes(4); !errors.Is(err, boom) {
		t.Fatalf("Bytes: %v", err)
	}
	if _, err := UUID(); !errors.Is(err, boom) {
		t.Fatalf("UUID: %v", err)
	}
	if _, err := UUIDWithLen(8); !errors.Is(err, boom) {
		t.Fatalf("UUIDWithLen: %v", err)
	}
	defer func() {
		if recover() == nil {
			t.Fatal("MustToken should panic on CSPRNG failure")
		}
	}()
	MustToken(8)
}

func TestShuffle(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8}
	Shuffle(in)
	sort.Ints(in)
	for i, v := range in {
		if v != i+1 {
			t.Fatalf("shuffle lost elements: %v", in)
		}
	}
	Shuffle([]int{})
}

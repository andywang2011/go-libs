package random_test

import (
	"slices"
	"testing"

	"github.com/andywang2011/go-libs/random"
)

func assertCharset(t *testing.T, got, charset string) {
	t.Helper()
	for _, r := range got {
		if !stringsContainsRune(charset, r) {
			t.Fatalf("char %q not in charset", r)
		}
	}
}

func stringsContainsRune(s string, r rune) bool {
	return slices.Contains([]rune(s), r)
}

func TestCharsLengthAndCharset(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(int) string
		charset string
		n       int
	}{
		{"Uppers", random.Uppers, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 10},
		{"Lowers", random.Lowers, "abcdefghijklmnopqrstuvwxyz", 10},
		{"Numbers", random.Numbers, "0123456789", 12},
		{"Strings", random.Strings, upperLowerNum(), 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn(tt.n)
			if len(got) != tt.n {
				t.Fatalf("len: got %d, want %d", len(got), tt.n)
			}
			assertCharset(t, got, tt.charset)
		})
	}
}

func upperLowerNum() string {
	return "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
}

func TestCharsEmpty(t *testing.T) {
	if random.Uppers(0) != "" {
		t.Fatal("expected empty string for n=0")
	}
}

func TestRangeNumInBounds(t *testing.T) {
	for i := 0; i < 50; i++ {
		n := random.RangeNum(3, 7)
		if n < 3 || n > 7 {
			t.Fatalf("RangeNum(3,7): %d", n)
		}
	}
}

func TestRangeNumSwappedBounds(t *testing.T) {
	for i := 0; i < 50; i++ {
		n := random.RangeNum(9, 4)
		if n < 4 || n > 9 {
			t.Fatalf("RangeNum(9,4): %d", n)
		}
	}
}

func TestRangeNumSingleton(t *testing.T) {
	if got := random.RangeNum(5, 5); got != 5 {
		t.Fatalf("RangeNum(5,5): %d", got)
	}
}

func TestShufflePermutation(t *testing.T) {
	orig := []int{1, 2, 3, 4, 5}
	sli := slices.Clone(orig)
	random.Shuffle(sli)
	if len(sli) != len(orig) {
		t.Fatal("length changed after shuffle")
	}
	slices.Sort(sli)
	if !slices.Equal(sli, orig) {
		t.Fatalf("elements changed: %v vs %v", sli, orig)
	}
}

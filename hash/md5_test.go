package hash_test

import (
	"testing"

	"github.com/andywang2011/go-libs/hash"
)

func TestMd5(t *testing.T) {
	got := hash.Md5("hello")
	want := "5d41402abc4b2a76b9719d911017c592"
	if got != want {
		t.Fatalf("Md5: got %q want %q", got, want)
	}
}

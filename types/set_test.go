package types_test

import (
	"testing"

	"github.com/andywang2011/go-libs/types"
)

func TestSet(t *testing.T) {
	s := types.NewSet([]string{"a", "b"})
	if s.Size() != 2 || !s.Has("a") || s.Has("c") {
		t.Fatalf("initial set: size=%d", s.Size())
	}
	s.Add("c")
	s.Remove("a")
	if s.Has("a") || !s.Has("c") {
		t.Fatal("Add/Remove failed")
	}
	items := s.Items()
	if len(items) != 2 {
		t.Fatalf("Items: %v", items)
	}
	s.Clear()
	if s.Size() != 0 {
		t.Fatal("Clear failed")
	}
}

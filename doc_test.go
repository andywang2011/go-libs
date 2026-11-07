package libs_test

import (
	"testing"

	"github.com/andywang2011/go-libs"
)

func TestPackageName(t *testing.T) {
	if libs.Doc() != "go-libs reserved module" {
		t.Fatalf("unexpected doc: %q", libs.Doc())
	}
}

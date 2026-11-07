package timex_test

import (
	"database/sql/driver"
	"encoding/json"
	"testing"
	"time"

	"github.com/andywang2011/go-libs/timex"
)

func TestJSONTimeMarshalUnmarshal(t *testing.T) {
	src := time.Date(2024, 6, 15, 10, 30, 45, 0, time.Local)
	jt := timex.JSONTime(src)

	b, err := json.Marshal(jt)
	if err != nil {
		t.Fatal(err)
	}
	want := `"2024-06-15 10:30:45"`
	if string(b) != want {
		t.Fatalf("marshal: got %s, want %s", b, want)
	}

	var out timex.JSONTime
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if time.Time(out) != src {
		t.Fatalf("unmarshal: got %v, want %v", time.Time(out), src)
	}
}

func TestJSONTimeUnmarshalInvalid(t *testing.T) {
	var jt timex.JSONTime
	err := json.Unmarshal([]byte(`"not-a-date"`), &jt)
	if err == nil {
		t.Fatal("expected parse error")
	}
}

func TestJSONTimeString(t *testing.T) {
	src := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)
	jt := timex.JSONTime(src)
	if jt.String() != "2020-01-02 03:04:05" {
		t.Fatalf("String: %q", jt.String())
	}
}

func TestJSONTimeValueAndScan(t *testing.T) {
	src := time.Date(2021, 7, 8, 9, 10, 11, 0, time.UTC)
	jt := timex.JSONTime(src)

	val, err := jt.Value()
	if err != nil {
		t.Fatal(err)
	}
	tm, ok := val.(time.Time)
	if !ok || !tm.Equal(src) {
		t.Fatalf("Value: %+v", val)
	}

	var scanned timex.JSONTime
	if err := scanned.Scan(src); err != nil {
		t.Fatal(err)
	}
	if time.Time(scanned) != src {
		t.Fatalf("Scan time: %v", scanned)
	}
}

func TestJSONTimeScanUnsupported(t *testing.T) {
	var jt timex.JSONTime
	err := jt.Scan("2024-01-01")
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "unsupported Scan type" {
		t.Fatalf("err: %v", err)
	}
}

func TestDateLayouts(t *testing.T) {
	if timex.DateTimeLayout != "2006-01-02 15:04:05" {
		t.Fatalf("DateTimeLayout: %q", timex.DateTimeLayout)
	}
	if timex.DateLayout != "2006-01-02" {
		t.Fatalf("DateLayout: %q", timex.DateLayout)
	}
	var _ driver.Valuer = timex.JSONTime{}
}

package jsonmap_test

import (
	"encoding/json"
	"testing"

	"github.com/andywang2011/go-libs/jsonmap"
)

func TestJSONMap_MarshalUnmarshal(t *testing.T) {
	m := jsonmap.JSONMap{"name": "Bob", "age": float64(20)}
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var decoded jsonmap.JSONMap
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded["name"] != "Bob" {
		t.Fatalf("decoded: %+v", decoded)
	}
}

func TestJSONMap_NilMarshal(t *testing.T) {
	var m jsonmap.JSONMap
	b, err := m.MarshalJSON()
	if err != nil || string(b) != "null" {
		t.Fatalf("nil marshal: %s err=%v", b, err)
	}
}

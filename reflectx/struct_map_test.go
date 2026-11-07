package reflectx_test

import (
	"errors"
	"testing"

	"github.com/andywang2011/go-libs/reflectx"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Desc     string
	OnlyOpt  string `json:",omitempty"`
	Ignore   string `json:"-"`
	DashName string `json:"-,"`
	secret   string
}

func TestStruct2Map(t *testing.T) {
	u := &User{ID: 123, Name: "Tom", Desc: "boy", OnlyOpt: "o", Ignore: "ignore", DashName: "d", secret: "xxx"}
	m, err := reflectx.Struct2Map(u)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]any{"id": 123, "name": "Tom", "Desc": "boy", "OnlyOpt": "o", "-": "d"}
	if len(m) != len(want) {
		t.Fatalf("got %v", m)
	}
	for k, v := range want {
		if m[k] != v {
			t.Errorf("key %q: got %v want %v", k, m[k], v)
		}
	}
	if _, ok := m["name,omitempty"]; ok {
		t.Error("omitempty must not leak into key")
	}
}

func TestStruct2MapValueAndDoublePointer(t *testing.T) {
	u := User{ID: 1}
	if m, err := reflectx.Struct2Map(u); err != nil || m["id"] != 1 {
		t.Fatalf("value: %v %v", m, err)
	}
	p := &u
	if m, err := reflectx.Struct2Map(&p); err != nil || m["id"] != 1 {
		t.Fatalf("ptr ptr: %v %v", m, err)
	}
}

func TestStruct2MapErrors(t *testing.T) {
	var nilUser *User
	for name, in := range map[string]any{
		"nil":         nil,
		"nil pointer": nilUser,
		"int":         1,
		"string":      "x",
		"slice":       []int{1},
		"map":         map[string]int{},
	} {
		m, err := reflectx.Struct2Map(in)
		if !errors.Is(err, reflectx.ErrNotStruct) || m != nil {
			t.Errorf("%s: got %v, %v", name, m, err)
		}
	}
}

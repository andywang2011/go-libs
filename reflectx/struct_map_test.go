package reflectx_test

import (
	"testing"

	"github.com/andywang2011/go-libs/reflectx"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string
	Ignore string `json:"-"`
	secret string
}

func TestStruct2Map(t *testing.T) {
	u := &User{
		ID:     123,
		Name:   "Tom",
		Desc:   "boy",
		Ignore: "ignore",
		secret: "xxx",
	}
	m := reflectx.Struct2Map(u)
	t.Log(m)
}

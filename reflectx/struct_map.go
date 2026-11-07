// Package reflectx provides small reflection helpers.
package reflectx

import (
	"errors"
	"reflect"
	"strings"
)

// ErrNotStruct is returned when Struct2Map receives a nil or non-struct value.
var ErrNotStruct = errors.New("reflectx: input must be a struct or non-nil pointer to struct")

// Struct2Map converts a struct (or pointer to struct) to a map keyed by the
// json tag name (the part before the first comma). Unexported fields and
// fields tagged `json:"-"` are skipped; fields without a name in the tag use
// the Go field name. Options such as omitempty do not affect the key.
func Struct2Map(obj any) (map[string]any, error) {
	if obj == nil {
		return nil, ErrNotStruct
	}
	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil, ErrNotStruct
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}
	t := v.Type()
	m := make(map[string]any, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() { // unexported fields panic on Interface()
			continue
		}
		tag := f.Tag.Get("json")
		if tag == "-" {
			continue
		}
		key, _, _ := strings.Cut(tag, ",")
		if key == "" {
			key = f.Name // no json name: use field name
		}
		m[key] = v.Field(i).Interface()
	}
	return m, nil
}

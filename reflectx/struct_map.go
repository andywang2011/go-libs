package reflectx

import (
	"reflect"
)

// Struct2Map converts a struct to a map using json struct tags as keys.
func Struct2Map(obj any) map[string]any {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	m := make(map[string]any)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() { // unexported fields panic on Interface()
			key := t.Field(i).Tag.Get("json")
			if key == "-" { // skip
				continue
			}
			if key == "" {
				key = t.Field(i).Name // no json tag: use field name
			}
			m[key] = v.Field(i).Interface()
		}
	}
	return m
}

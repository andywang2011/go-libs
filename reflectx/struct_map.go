package reflectx

import (
	"reflect"
)

// Struct2Map 结构体转成map
func Struct2Map(obj any) map[string]any {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	m := make(map[string]any)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() { // 不可导出字段取Interface会panic
			key := t.Field(i).Tag.Get("json")
			if key == "-" { // 忽略
				continue
			}
			if key == "" {
				key = t.Field(i).Name // 未指定json tag 的情况，使用字段名
			}
			m[key] = v.Field(i).Interface()
		}
	}
	return m
}

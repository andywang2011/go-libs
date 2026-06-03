package types

func SliceEqual[T comparable](a, b []T) bool {
	la, lb := len(a), len(b)
	if la != lb {
		return false
	}
	//return reflect.DeepEqual(a, b) // 反射效率较差，改用遍历比较
	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func SliceToMap[T comparable](arr []T) map[T]struct{} {
	m := make(map[T]struct{}, len(arr))
	for _, v := range arr {
		m[v] = struct{}{}
	}
	return m
}

func SliceToCountMap[T comparable](arr []T) map[T]int {
	m := make(map[T]int, len(arr))
	for _, v := range arr {
		m[v]++
	}
	return m
}

/*
Go1.21标准库将新增slices包和maps包提供对任意类型slice和map的常用操作，不再需要以下方法。
*/

func SliceContain[T comparable](arr []T, target T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[M ~map[K]V, K comparable, V any](m M) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

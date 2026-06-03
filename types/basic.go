package types

/*
Go 1.21 新增标准库cmp中定义了Ordered泛型
*/

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

/*
获取一组可比较大小类型里面的最值，Go 1.21 以后新增内置的 min 和 max 函数实现了以下两个等效方法。
*/

func Max[T Ordered](item T, items ...T) T {
	max := item
	for _, v := range items {
		if max < v {
			max = v
		}
	}
	return max
}

func Min[T Ordered](item T, items ...T) T {
	min := item
	for _, v := range items {
		if min > v {
			min = v
		}
	}
	return min
}

package types

/*
Go 1.21 stdlib cmp defines the Ordered constraint; this package mirrors it for pre-1.21 builds.
*/

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

/*
Max and Min mirror the built-in min/max added in Go 1.21 for ordered types.
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

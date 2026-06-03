package types

// Set 非并发安全的简单集合
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet 初始化一个集合，必须传入类型明确的slice
func NewSet[T comparable](items []T) *Set[T] {
	m := make(map[T]struct{})
	for _, item := range items {
		m[item] = struct{}{}
	}
	return &Set[T]{m: m}
}

// Add 往集合添加元素
func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

// Remove 从集合移除元素
func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}

// Clear 清空集合
func (s *Set[T]) Clear() {
	s.m = map[T]struct{}{}
}

// Size 集合元素个数
func (s *Set[T]) Size() int {
	return len(s.m)
}

// Has 集合中是否包含某元素
func (s *Set[T]) Has(v T) bool {
	_, ok := s.m[v]
	return ok
}

// Items 集合元素
func (s *Set[T]) Items() []T {
	items := make([]T, 0, len(s.m))
	for item := range s.m {
		items = append(items, item)
	}
	return items
}

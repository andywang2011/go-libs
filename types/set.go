package types

// Set is a simple, non-concurrent-safe set.
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet builds a set from a typed slice.
func NewSet[T comparable](items []T) *Set[T] {
	m := make(map[T]struct{})
	for _, item := range items {
		m[item] = struct{}{}
	}
	return &Set[T]{m: m}
}

// Add inserts an element.
func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

// Remove deletes an element.
func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}

// Clear removes all elements.
func (s *Set[T]) Clear() {
	s.m = map[T]struct{}{}
}

// Size returns the number of elements.
func (s *Set[T]) Size() int {
	return len(s.m)
}

// Has reports whether v is in the set.
func (s *Set[T]) Has(v T) bool {
	_, ok := s.m[v]
	return ok
}

// Items returns all elements (order undefined).
func (s *Set[T]) Items() []T {
	items := make([]T, 0, len(s.m))
	for item := range s.m {
		items = append(items, item)
	}
	return items
}

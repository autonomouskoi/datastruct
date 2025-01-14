// Package mapset provides a set implementation using map[T]struct{}
package mapset

// MapSet is a set of T
type MapSet[T comparable] map[T]struct{}

// From creates a new set initialized with the provided values.
func From[T comparable](vs ...T) MapSet[T] {
	ms := make(MapSet[T], len(vs))
	for _, v := range vs {
		ms.Add(v)
	}
	return ms
}

// Add a value to the set
func (ms MapSet[T]) Add(v T) {
	ms[v] = struct{}{}
}

// Equals returns whether or not two sets have the same items
func (ms MapSet[T]) Equals(o MapSet[T]) bool {
	if len(ms) != len(o) {
		return false
	}
	for v := range ms {
		if !o.Has(v) {
			return false
		}
	}
	return true
}

// Has returns whether the set has the item
func (ms MapSet[T]) Has(v T) bool {
	_, present := ms[v]
	return present
}

// Slice returns a slice of the set's items in no particular order
func (ms MapSet[T]) Slice() []T {
	s := make([]T, 0, len(ms))
	for k := range ms {
		s = append(s, k)
	}
	return s
}

// Subtract updates ms to not include members of o
func (ms MapSet[T]) Subtract(o MapSet[T]) {
	for k := range o {
		delete(ms, k)
	}
}

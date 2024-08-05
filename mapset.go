package mapset

type MapSet[T comparable] map[T]struct{}

func From[T comparable](vs ...T) MapSet[T] {
	ms := make(MapSet[T], len(vs))
	for _, v := range vs {
		ms.Add(v)
	}
	return ms
}

func (ms MapSet[T]) Add(v T) {
	ms[v] = struct{}{}
}

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

func (ms MapSet[T]) Has(v T) bool {
	_, present := ms[v]
	return present
}

func (ms MapSet[T]) Slice() []T {
	s := make([]T, 0, len(ms))
	for k := range ms {
		s = append(s, k)
	}
	return s
}

func (ms MapSet[T]) Subtract(o MapSet[T]) {
	for k := range o {
		delete(ms, k)
	}
}

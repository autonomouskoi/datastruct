package iterutil

import "iter"

// Map returns a new iterator that applies function fn to each value produced
// by iterator i.
func Map[I, O any](i iter.Seq[I], fn func(I) O) iter.Seq[O] {
	return func(yield func(O) bool) {
		for iv := range i {
			if !yield(fn(iv)) {
				return
			}
		}
	}
}

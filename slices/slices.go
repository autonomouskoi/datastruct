package slices

func Map[T, U any](in []T, fn func(T) U) []U {
	out := make([]U, len(in))
	for i, t := range in {
		out[i] = fn(t)
	}
	return out
}

func Matches[T any](in []T, fn func(T) bool) bool {
	for _, t := range in {
		if fn(t) {
			return true
		}
	}
	return false
}

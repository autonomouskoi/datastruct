package slices

func Map[T, U any](in []T, fn func(T) U) []U {
	out := make([]U, len(in))
	for i, t := range in {
		out[i] = fn(t)
	}
	return out
}

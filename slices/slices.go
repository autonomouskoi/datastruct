// Package slices provide utility functions for slices
package slices

// Map applies a function to a slice, returning a new slice of values returned
// by fn
func Map[T, U any](in []T, fn func(T) U) []U {
	out := make([]U, len(in))
	for i, t := range in {
		out[i] = fn(t)
	}
	return out
}

// Matches calls fn on each value in in. Matches returns true the first time
// fn returns true. It returns false if no member of in causes fn to return true
func Matches[T any](in []T, fn func(T) bool) bool {
	for _, t := range in {
		if fn(t) {
			return true
		}
	}
	return false
}

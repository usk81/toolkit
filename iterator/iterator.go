package iterator

type (
	// Iterator ...
	Iterator[T any] interface {
		Next() bool
		Value() T
	}

	// Reducer ...
	Reducer[T, V any] func(accum T, value V) T
)

// Collect ...
func Collect[T any](iter Iterator[T]) []T {
	var xs []T
	for iter.Next() {
		xs = append(xs, iter.Value())
	}
	return xs
}

// Reduce values iterated over to a single value
func Reduce[T, V any](iter Iterator[V], f Reducer[T, V]) T {
	var accum T
	for iter.Next() {
		accum = f(accum, iter.Value())
	}
	return accum
}

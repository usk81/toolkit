package iterator

type (
	// FilterIterator ...
	FilterIterator[T any] struct {
		source Iterator[T]
		pred   func(T) bool
	}
)

// Filter keeps only those values which match a supplied value
func Filter[T any](iter Iterator[T], pred func(T) bool) Iterator[T] {
	return &FilterIterator[T]{
		iter, pred,
	}
}

// Next ...
func (f *FilterIterator[T]) Next() bool {
	for f.source.Next() {
		if f.pred(f.source.Value()) {
			return true
		}
	}
	return false
}

// Value ...
func (f *FilterIterator[T]) Value() T {
	return f.source.Value()
}

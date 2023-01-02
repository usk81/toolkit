package iterator

type (
	// SliceIterator ...
	SliceIterator[T any] struct {
		Elements []T
		value    T
		index    int
	}
)

// Slice creates an iterator over the slice xs
func Slice[T any](xs []T) Iterator[T] {
	return &SliceIterator[T]{
		Elements: xs,
	}
}

// Next moves to next value in collection
func (s *SliceIterator[T]) Next() bool {
	if s.index < len(s.Elements) {
		s.value = s.Elements[s.index]
		s.index += 1
		return true
	}

	return false
}

// Value gets current element
func (s *SliceIterator[T]) Value() T {
	return s.value
}

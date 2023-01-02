package iterator

type (
	// Map ...
	MapIterator[T any] struct {
		source Iterator[T]
		mapper func(T) T
	}
)

// Next ...
func (m *MapIterator[T]) Next() bool {
	return m.source.Next()
}

// Value ...
func (m *MapIterator[T]) Value() T {
	value := m.source.Value()
	return m.mapper(value)
}

// Map ...
func Map[T any](iter Iterator[T], f func(T) T) Iterator[T] {
	return &MapIterator[T]{
		iter, f,
	}
}

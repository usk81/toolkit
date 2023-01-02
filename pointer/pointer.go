package pointer

// IsNil checks if value is nil
func IsNil(v interface{}) bool {
	return v == nil
}

// To converts to pointer
func To[T any](v T) *T {
	if IsNil(v) {
		return nil
	}
	return &v
}

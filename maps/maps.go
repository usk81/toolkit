package maps

import "errors"

// Combine creates an slice by using one slice for keys and another for its values
func Combine[K comparable, V any](keys []K, values []V) (map[K]V, error) {
	if len(keys) != len(values) {
		return nil, errors.New("argument #1 (keys) and argument #2 (values) must have the same number of elements")
	}
	result := map[K]V{}
	for i := range keys {
		result[keys[i]] = values[i]
	}
	return result, nil
}

// Keys returns the list of all keys of the given map
func Keys[K comparable, V any](m map[K]V) []K {
	if len(m) == 0 {
		return nil
	}
	ks := make([]K, len(m))
	i := 0
	for k := range m {
		ks[i] = k
		i++
	}
	return ks
}

package slice

import (
	"errors"
	"fmt"
	"reflect"
)

// Chunk creates an array of elements split into groups the length of size.
// If array can't be split evenly, the final chunk will be the remaining elements
func Chunk[T any](vs []T, size int) (rs [][]T, err error) {
	if vs == nil {
		return
	}
	if size <= 0 {
		err = errors.New("size must be greater then 0")
		return
	}

	n := len(vs) / size
	if len(vs)%size != 0 {
		n += 1
	}

	rs = make([][]T, n)
	for i := 0; i < n; i++ {
		last := (i + 1) * size
		if last > len(vs) {
			last = len(vs)
		}
		rs[i] = vs[i*size : last]
	}
	return
}

// Compact creates an slice with all zero values removed
func Compact[T comparable](vs []T) []T {
	var zero T

	rs := []T{}
	for _, v := range vs {
		if v != zero {
			rs = append(rs, v)
		}
	}
	return rs
}

func Difference[T comparable](vs, xs []T) []T {
	if len(vs) == 0 || len(xs) == 0 {
		return vs
	}
	dropper := func(v T) bool {
		return Exists(v, xs)
	}
	return Drop(vs, dropper)
}

// Drop creates a slice excluding some elements dropped. if dropper returns true, this element is removed.
func Drop[T comparable](vs []T, dropper func(v T) bool) []T {
	result := []T{}
	for _, v := range vs {
		if !dropper(v) {
			result = append(result, v)
		}
	}
	return result
}

// Flattens slice a single level deep
func Flatten[T any](v interface{}) ([]T, error) {
	if v == nil {
		return nil, nil
	}
	return doFlatten([]T{}, v)
}

func doFlatten[T any](ss []T, x interface{}) ([]T, error) {
	var err error
	switch v := x.(type) {
	case T:
		ss = append(ss, v)
	case []T:
		ss = append(ss, v...)
	case []interface{}:
		for i := range v {
			if ss, err = doFlatten(ss, v[i]); err != nil {
				return nil, err
			}
		}
	default:
		rv := reflect.ValueOf(x)
		if rv.Kind() == reflect.Slice {
			for i := 0; i < rv.Len(); i++ {
				if ss, err = doFlatten(ss, rv.Index(i).Interface()); err != nil {
					return nil, err
				}
			}
		} else {
			return nil, fmt.Errorf("not vaild value given. type: %v", reflect.ValueOf(x).Kind())
		}
	}
	return ss, nil
}

// Exists checks if a value exists in an slice
func Exists[T comparable](v T, xs []T) bool {
	for _, x := range xs {
		if v == x {
			return true
		}
	}
	return false
}

func Intersection[T comparable](vs, xs []T) []T {
	if len(vs) == 0 || len(xs) == 0 {
		return []T{}
	}
	dropper := func(v T) bool {
		return !Exists(v, xs)
	}
	return Drop(vs, dropper)
}

func Reduce[T, R any](vs []T, f func(acc R, v T, index int) R, initial R) R {
	acc := initial
	for i := range vs {
		acc = f(acc, vs[i], i)
	}
	return acc
}

// Unique Removes duplicate values from an slice
func Unique[T comparable](xs []T) []T {
	if xs == nil {
		return nil
	}
	m := map[T]struct{}{}
	for _, x := range xs {
		m[x] = struct{}{}
	}
	u := make([]T, len(m))
	i := 0
	for k := range m {
		u[i] = k
		i++
	}
	return u
}

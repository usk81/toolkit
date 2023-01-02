package testkit

import "testing"

type Testable interface {
	Test(t *testing.T)
}

type TestCase[C any] struct {
	Describe string
	Cases    []C
	Runner   func(t *testing.T, c C)
}

func (tc TestCase[C]) Test(t *testing.T) {
	t.Run(tc.Describe, func(t *testing.T) {
		for _, tt := range tc.Cases {
			// for Parallel tests
			tt := tt
			tc.Runner(t, tt)
		}
	})
}

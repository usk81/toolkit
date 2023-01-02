package pointer

import (
	"testing"

	"github.com/usk81/toolkit/testkit"
)

func TestTo(t *testing.T) {
	type args[T any] struct {
		v T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *T
	}

	tests := []testkit.Testable{}

	// string
	tests = append(tests, testkit.TestCase[testCase[string]]{
		Describe: "case_string",
		Cases: []testCase[string]{
			{
				name: "some_string",
				args: args[string]{
					v: "foo",
				},
				want: func() *string {
					v := "foo"
					return &v
				}(),
			},
			{
				name: "empty",
				args: args[string]{
					v: "",
				},
				want: func() *string {
					v := ""
					return &v
				}(),
			},
		},
		Runner: func(t *testing.T, tt testCase[string]) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				if got := To(tt.args.v); *got != *tt.want {
					t.Errorf("To() = %s, want %s", *got, *tt.want)
				}
			})
		},
	})

	// int
	tests = append(tests, testkit.TestCase[testCase[int]]{
		Describe: "case_int",
		Cases: []testCase[int]{
			{
				name: "some_integer",
				args: args[int]{
					v: 1234567890,
				},
				want: func() *int {
					v := 1234567890
					return &v
				}(),
			},
			{
				name: "zero",
				args: args[int]{
					v: 0,
				},
				want: func() *int {
					v := 0
					return &v
				}(),
			},
		},
		Runner: func(t *testing.T, tt testCase[int]) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				if got := To(tt.args.v); *got != *tt.want {
					t.Errorf("To() = %d, want %d", *got, *tt.want)
				}
			})
		},
	})

	// bool
	tests = append(tests, testkit.TestCase[testCase[bool]]{
		Describe: "case_bool",
		Cases: []testCase[bool]{
			{
				name: "true",
				args: args[bool]{
					v: true,
				},
				want: func() *bool {
					v := true
					return &v
				}(),
			},
			{
				name: "false",
				args: args[bool]{
					v: false,
				},
				want: func() *bool {
					v := false
					return &v
				}(),
			},
		},
		Runner: func(t *testing.T, tt testCase[bool]) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				if got := To(tt.args.v); *got != *tt.want {
					t.Errorf("To() = %t, want %t", *got, *tt.want)
				}
			})
		},
	})

	// any
	tests = append(tests, testkit.TestCase[testCase[any]]{
		Describe: "case_any",
		Cases: []testCase[any]{
			{
				name: "nil",
				args: args[any]{
					v: nil,
				},
				want: nil,
			},
		},
		Runner: func(t *testing.T, tt testCase[any]) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				if got := To(tt.args.v); got != tt.want {
					t.Errorf("To() = %s, want %s", *got, *tt.want)
				}
			})
		},
	})

	for _, tt := range tests {
		tt := tt
		tt.Test(t)
	}
}

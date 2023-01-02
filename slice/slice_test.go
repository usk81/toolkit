package slice

import (
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/usk81/toolkit/testkit"
	"golang.org/x/exp/constraints"
)

func less[T constraints.Ordered](a, b T) bool { return a < b }

func lessBool(_, b bool) bool { return b }

func TestExists(t *testing.T) {
	type args[T comparable] struct {
		v  T
		xs []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}

	tests := []testkit.Testable{}

	// string
	tests = append(tests, testkit.TestCase[testCase[string]]{
		Describe: "case_string",
		Cases: []testCase[string]{
			{
				name: "exist",
				args: args[string]{
					v: "foo",
					xs: []string{
						"foo",
						"bar",
					},
				},
				want: true,
			},
			{
				name: "not_exist",
				args: args[string]{
					v: "fizz",
					xs: []string{
						"foo",
						"bar",
					},
				},
				want: false,
			},
			{
				name: "empty_value",
				args: args[string]{
					v: "",
					xs: []string{
						"",
						"empty",
					},
				},
				want: true,
			},
			{
				name: "empty_slice",
				args: args[string]{
					v:  "foo",
					xs: []string{},
				},
				want: false,
			},
			{
				name: "nil",
				args: args[string]{
					v:  "foo",
					xs: nil,
				},
				want: false,
			},
		},
		Runner: func(t *testing.T, tt testCase[string]) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				if got := Exists(tt.args.v, tt.args.xs); got != tt.want {
					t.Errorf("Exists() = %v, want %v", got, tt.want)
				}
			})
		},
	})

	// int
	tests = append(tests, testkit.TestCase[testCase[int]]{
		Describe: "case_int",
		Cases: []testCase[int]{
			{
				name: "exist",
				args: args[int]{
					v: 123,
					xs: []int{
						123,
						456,
					},
				},
				want: true,
			},
			{
				name: "not_exist",
				args: args[int]{
					v: 789,
					xs: []int{
						123,
						456,
					},
				},
				want: false,
			},
			{
				name: "zero",
				args: args[int]{
					v: 0,
					xs: []int{
						0,
						1,
					},
				},
				want: true,
			},
			{
				name: "max_int32",
				args: args[int]{
					v: math.MaxInt32,
					xs: []int{
						math.MaxInt32,
					},
				},
				want: true,
			},
			{
				name: "max_int64",
				args: args[int]{
					v: math.MaxInt64,
					xs: []int{
						math.MaxInt64,
					},
				},
				want: true,
			},
			{
				name: "empty_slice",
				args: args[int]{
					v:  0,
					xs: []int{},
				},
				want: false,
			},
			{
				name: "nil",
				args: args[int]{
					v:  0,
					xs: nil,
				},
				want: false,
			},
		},
		Runner: func(t *testing.T, tt testCase[int]) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				if got := Exists(tt.args.v, tt.args.xs); got != tt.want {
					t.Errorf("Exists() = %v, want %v", got, tt.want)
				}
			})
		},
	})

	// boolean
	tests = append(tests, testkit.TestCase[testCase[bool]]{
		Describe: "case_bool",
		Cases: []testCase[bool]{
			{
				name: "exist_true",
				args: args[bool]{
					v: true,
					xs: []bool{
						true,
						false,
					},
				},
				want: true,
			},
			{
				name: "exist_false",
				args: args[bool]{
					v: false,
					xs: []bool{
						true,
						false,
					},
				},
				want: true,
			},
			{
				name: "not_exist",
				args: args[bool]{
					v: true,
					xs: []bool{
						false,
					},
				},
				want: false,
			},
			{
				name: "empty_slice",
				args: args[bool]{
					v:  false,
					xs: []bool{},
				},
				want: false,
			},
			{
				name: "nil",
				args: args[bool]{
					v:  false,
					xs: nil,
				},
				want: false,
			},
		},
		Runner: func(t *testing.T, tt testCase[bool]) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				if got := Exists(tt.args.v, tt.args.xs); got != tt.want {
					t.Errorf("Exists() = %v, want %v", got, tt.want)
				}
			})
		},
	})

	for _, tt := range tests {
		tt := tt
		tt.Test(t)
	}
}

func TestUnique(t *testing.T) {
	type args[T comparable] struct {
		xs []T
	}
	type testCaseForUnique[T comparable] struct {
		name string
		args args[T]
		want []T
	}

	tests := []testkit.Testable{}

	// string
	tests = append(tests, testkit.TestCase[testCaseForUnique[string]]{
		Describe: "case_string",
		Cases: []testCaseForUnique[string]{
			{
				name: "unique",
				args: args[string]{
					xs: []string{
						"foo",
						"bar",
					},
				},
				want: []string{
					"foo",
					"bar",
				},
			},
			{
				name: "duplicated",
				args: args[string]{
					xs: []string{
						"foo",
						"bar",
						"foo",
					},
				},
				want: []string{
					"foo",
					"bar",
				},
			},
			{
				name: "have_empty_value",
				args: args[string]{
					xs: []string{
						"foo",
						"",
					},
				},
				want: []string{
					"foo",
					"",
				},
			},
			{
				name: "empty_slice",
				args: args[string]{
					xs: []string{},
				},
				want: []string{},
			},
		},
		Runner: func(t *testing.T, tt testCaseForUnique[string]) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Unique(tt.args.xs)
				if diff := cmp.Diff(got, tt.want, cmpopts.SortSlices(less[string])); diff != "" {
					t.Errorf("Unique() = %v, want %v", got, tt.want)
				}
			})
		},
	})

	// int
	tests = append(tests, testkit.TestCase[testCaseForUnique[int]]{
		Describe: "case_int",
		Cases: []testCaseForUnique[int]{
			{
				name: "unique",
				args: args[int]{
					xs: []int{
						123,
						456,
					},
				},
				want: []int{
					123,
					456,
				},
			},
			{
				name: "duplicated",
				args: args[int]{
					xs: []int{
						123,
						456,
						456,
					},
				},
				want: []int{
					123,
					456,
				},
			},
			{
				name: "have_zero",
				args: args[int]{
					xs: []int{
						1,
						0,
					},
				},
				want: []int{
					1,
					0,
				},
			},
			{
				name: "empty_slice",
				args: args[int]{
					xs: []int{},
				},
				want: []int{},
			},
		},
		Runner: func(t *testing.T, tt testCaseForUnique[int]) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Unique(tt.args.xs)
				if diff := cmp.Diff(got, tt.want, cmpopts.SortSlices(less[int])); diff != "" {
					t.Errorf("Unique() = %v, want %v", got, tt.want)
				}
			})
		},
	})

	// float
	tests = append(tests, testkit.TestCase[testCaseForUnique[float64]]{
		Describe: "case_float64",
		Cases: []testCaseForUnique[float64]{
			{
				name: "unique",
				args: args[float64]{
					xs: []float64{
						1.05,
						1.08,
					},
				},
				want: []float64{
					1.05,
					1.08,
				},
			},
			{
				name: "duplicated",
				args: args[float64]{
					xs: []float64{
						1.05,
						1.08,
						1.08,
					},
				},
				want: []float64{
					1.05,
					1.08,
				},
			},
			{
				name: "have_zero",
				args: args[float64]{
					xs: []float64{
						0,
						3.14,
					},
				},
				want: []float64{
					0,
					3.14,
				},
			},
			{
				name: "empty_slice",
				args: args[float64]{
					xs: []float64{},
				},
				want: []float64{},
			},
		},
		Runner: func(t *testing.T, tt testCaseForUnique[float64]) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Unique(tt.args.xs)
				if diff := cmp.Diff(got, tt.want, cmpopts.SortSlices(less[float64])); diff != "" {
					t.Errorf("Unique() = %v, want %v", got, tt.want)
				}
			})
		},
	})

	// boolean
	tests = append(tests, testkit.TestCase[testCaseForUnique[bool]]{
		Describe: "case_bool",
		Cases: []testCaseForUnique[bool]{
			{
				name: "unique",
				args: args[bool]{
					xs: []bool{
						true,
						false,
					},
				},
				want: []bool{
					true,
					false,
				},
			},
			{
				name: "duplicated",
				args: args[bool]{
					xs: []bool{
						true,
						false,
						false,
					},
				},
				want: []bool{
					true,
					false,
				},
			},
			{
				name: "empty_slice",
				args: args[bool]{
					xs: []bool{},
				},
				want: []bool{},
			},
		},
		Runner: func(t *testing.T, tt testCaseForUnique[bool]) {
			t.Run(tt.name, func(t *testing.T) {
				got := Unique(tt.args.xs)
				if diff := cmp.Diff(got, tt.want, cmpopts.SortSlices(lessBool)); diff != "" {
					t.Errorf("Unique() = %v, want %v", got, tt.want)
				}
			})
		},
	})

	for _, tt := range tests {
		tt := tt
		tt.Test(t)
	}
}

func TestFlatten(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "a",
			args: args{
				v: [][][]string{
					{
						{
							"foobar",
						},
					},
				},
			},
			want: []string{
				"foobar",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Flatten[string](tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Flatten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

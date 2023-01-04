package maps

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/usk81/toolkit/testkit"
	"golang.org/x/exp/constraints"
)

func less[K constraints.Ordered](a, b K) bool { return a < b }

func TestCombine(t *testing.T) {
	type args[K comparable] struct {
		keys   []K
		values []any
	}
	type testCaseForCombine[K comparable] struct {
		name    string
		args    args[K]
		want    map[K]any
		wantErr bool
	}

	tests := []testkit.Testable{}

	tests = append(tests, testkit.TestCase[testCaseForCombine[string]]{
		Describe: "case_string",
		Cases: []testCaseForCombine[string]{
			{
				name: "same number of elements",
				args: args[string]{
					keys: []string{
						"foo",
						"bar",
					},
					values: []interface{}{
						584670,
						104860,
					},
				},
				want: map[string]interface{}{
					"foo": 584670,
					"bar": 104860,
				},
				wantErr: false,
			},
			{
				name: "not same number of elements",
				args: args[string]{
					keys: []string{
						"a",
						"b",
					},
					values: []interface{}{
						"z",
						"y",
						"x",
					},
				},
				want:    nil,
				wantErr: true,
			},
			{
				name: "include same keys",
				args: args[string]{
					keys: []string{
						"a",
						"b",
						"c",
						"b",
					},
					values: []interface{}{
						"z",
						"y",
						"x",
						"w",
					},
				},
				want: map[string]interface{}{
					"a": "z",
					"b": "w",
					"c": "x",
				},
				wantErr: false,
			},
		},
		Runner: func(t *testing.T, tt testCaseForCombine[string]) {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got, err := Combine(tt.args.keys, tt.args.values)
				if (err != nil) != tt.wantErr {
					t.Errorf("Combine() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if diff := cmp.Diff(got, tt.want, cmpopts.SortMaps(less[string])); diff != "" {
					t.Errorf("Combine() = %v, want %v, diff %s", got, tt.want, diff)
				}
			})
		},
	})

	for _, tt := range tests {
		tt := tt
		tt.Test(t)
	}
}

func TestKeys(t *testing.T) {
	type args[K comparable] struct {
		values map[K]interface{}
	}

	type testCaseForKeys[K comparable] struct {
		name string
		args args[K]
		want []K
	}

	tests := []testkit.Testable{}

	tests = append(tests, testkit.TestCase[testCaseForKeys[string]]{
		Describe: "case_string",
		Cases: []testCaseForKeys[string]{
			{
				name: "return slice",
				args: args[string]{
					values: map[string]interface{}{
						"Mercury": "水星",
						"Venus":   "金星",
						"Earth":   "地球",
						"Mars":    "火星",
						"Jupiter": "木星",
						"Saturn":  "土星",
						"Uranus":  "天王星",
						"Neptune": "海王星",
					},
				},
				want: []string{
					"Mercury",
					"Venus",
					"Earth",
					"Mars",
					"Jupiter",
					"Saturn",
					"Uranus",
					"Neptune",
				},
			},
			{
				name: "empty map",
				args: args[string]{
					values: map[string]interface{}{},
				},
				want: nil,
			},
			{
				name: "nil",
				args: args[string]{
					values: nil,
				},
				want: nil,
			},
		},
		Runner: func(t *testing.T, tt testCaseForKeys[string]) {
			t.Parallel()
			t.Run(tt.name, func(t *testing.T) {
				got := Keys(tt.args.values)
				if diff := cmp.Diff(got, tt.want, cmpopts.SortSlices(less[string])); diff != "" {
					t.Errorf("Combine() = %v, want %v, diff %s", got, tt.want, diff)
				}
			})
		},
	})

	for _, tt := range tests {
		tt := tt
		tt.Test(t)
	}
}

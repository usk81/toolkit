package maps

import (
	"fmt"
	"sort"
)

func ExampleCombine() {
	keys := []string{"green", "red", "yellow"}
	values := []string{"avocado", "apple", "banana"}
	result, _ := Combine(keys, values)

	fmt.Println(result)
	// Output: map[green:avocado red:apple yellow:banana]
}

func ExampleKeys() {
	vs := map[string]int{
		"foo":  123,
		"bar":  456,
		"fizz": 789,
	}

	keys := Keys(vs)
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	fmt.Println(keys)
	// Output: [bar fizz foo]
}

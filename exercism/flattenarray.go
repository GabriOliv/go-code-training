package main

import (
	"fmt"
)

// func flattenarray(array []int) []int {
// 	fmt.Println(array)

// 	return nil
// }

func flattenArray(array interface{}) []interface{} {
	var output []interface{}

	switch item := array.(type) {
	case nil:
	case []interface{}:
		for _, nested := range item {
			output = append(output, flattenArray(nested)...)
		}
	default:
		output = append(output, item)
	}

	return output
}

func main() {

	var input []interface{}

	input = append(input,
		1,
		append(input, 2, 3, nil, 4),
		append(input, nil),
		5,
	)

	// [1,[2,3,null,4],[null],5]
	fmt.Println(input)
	// [1 2 3 4 5]
	fmt.Println(flattenArray(input))

}

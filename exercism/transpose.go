package main

import (
	"fmt"
	"strings"
)

func transpose(text string) string {
	// string to []string
	matrix := strings.Split(text, string(rune(10)))

	max_length := maxlength(matrix)
	transposed := make([]string, max_length)

	// WhiteSpace
	for index, line := range matrix {
		next_max_length := maxlength(matrix[index+1:])
		row_length := len(line)

		if (next_max_length - row_length) > 0 {
			matrix[index] += strings.Repeat(" ", next_max_length-row_length)
		}
	}

	// Transpose
	for _, row := range matrix {
		for j, col := range row {
			transposed[j] += string(col)
		}
	}

	// []string to string
	text = strings.Join(transposed, string(rune(10)))

	return text
}

func maxlength(matrix []string) int {
	length := 0
	for _, row := range matrix {
		if len(row) > length {
			length = len(row)
		}
	}
	return length
}

func main() {

	str := `This is a
multiline string.
Transposed with golang.
a small line,
and one big line.

ABC
DEF

AB
DEF
`

	fmt.Println(str)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	str = transpose(str)

	fmt.Println(str)
}

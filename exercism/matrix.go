package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func convertToMatrix(text string) ([][]int, error) {

	// Match string with numbers
	if regexp.MustCompile(`^(\d+( |\n|$)*)*$`).MatchString(text) == false {
		return nil, errors.New("invalid input")
	}

	var matrix [][]int

	// String to matrix of int
	for _, lines := range strings.Split(text, string(rune(10))) {
		aux := []int{}
		for _, elem := range strings.Split(lines, " ") {
			if s, err := strconv.Atoi(elem); err == nil {
				aux = append(aux, s)
			}
		}
		matrix = append(matrix, aux)
	}

	// Fill non-square matrix
	for key := range matrix {
		for len(matrix[key]) < maxLength(matrix) {
			matrix[key] = append(matrix[key], 0)
		}
	}

	return matrix, nil
}

func list_rows(matrix [][]int) {
	for _, row := range matrix {
		for _, elem := range row {
			fmt.Print(elem, ", ")
		}
		fmt.Println()
	}
}

func list_cols(matrix [][]int) {
	rows := len(matrix)
	cols := maxLength(matrix)

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			fmt.Print(matrix[j][i], ", ")
		}
		fmt.Println()
	}
}

func maxLength(matrix [][]int) int {
	length := 0
	for _, row := range matrix {
		if len(row) > length {
			length = len(row)
		}
	}
	return length
}

func replaceValue(matrix [][]int, row int, col int, number int) ([][]int, error) {
	if (row < 0 || row >= len(matrix)) || (col < 0 || col >= maxLength(matrix)) {
		return nil, errors.New("invalid position")
	}

	matrix[row][col] = number

	return matrix, nil
}

func main() {
	text := `9 8 7
5 3
6 6 7 8 9 `

	fmt.Println("\nText:")
	fmt.Println(text)

	matrix, err := convertToMatrix(text)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nMatrix:")
	fmt.Println(matrix)

	fmt.Println("\nRows:")
	list_rows(matrix)

	fmt.Println("\nColumns:")
	list_cols(matrix)

	matrix, err = replaceValue(matrix, 0, 4, 42)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	list_rows(matrix)

}

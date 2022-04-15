package main

import (
	"errors"
	"fmt"
)

func spiralMatrix(size int) ([][]int, error) {
	// Error with non-positive numbers
	if size < 1 {
		return nil, errors.New("invalid size")
	}

	// Declare slice
	matrix := make([][]int, size, size)
	// Fill lines
	for key := range matrix {
		matrix[key] = make([]int, size, size)
	}

	// Limits of x move
	cap_x := size
	low_x := 0
	// Limits of y move
	cap_y := size
	low_y := 0
	// Number inside cell
	counter := 1

	// Count moves
	for m := 0; m < (2*size - 1); m++ {
		// Move Right
		if m%4 == 0 {
			for i := low_x; i < cap_x; i++ {
				matrix[low_x][i] = counter
				counter++
			}
			low_y++
		}
		// Move Down
		if m%4 == 1 {
			for i := low_y; i < cap_y; i++ {
				matrix[i][cap_x-1] = counter
				counter++
			}
			cap_x--
		}
		// Move Left
		if m%4 == 2 {
			for i := cap_x - 1; i >= low_x; i-- {
				matrix[cap_y-1][i] = counter
				counter++
			}
			cap_y--
		}
		// Move Up
		if m%4 == 3 {
			for i := cap_y - 1; i >= low_y; i-- {
				matrix[i][low_y-1] = counter
				counter++
			}
			low_x++
		}
	}

	return matrix, nil
}

func main() {

	fmt.Println()

	for key := 1; key < 7; key++ {

		matrix, err := spiralMatrix(key)

		if err != nil {
			panic(err)
		}

		for _, row := range matrix {
			for _, elem := range row {
				fmt.Printf("%02d ", elem)
			}
			fmt.Println()
		}

		fmt.Println()
	}

}

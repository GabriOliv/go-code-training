package main

import "fmt"

func chessboard(grains uint64, squares int) []uint64 {
	var board []uint64

	for squares > 0 {
		board = append(board, grains)
		grains <<= 1
		squares--
	}

	return board
}

func main() {

	grains := 1

	board := chessboard(uint64(grains), 64)

	fmt.Println(board)

}

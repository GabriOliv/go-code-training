package main

import (
	"fmt"
)

func reversestring(phrase string) string {
	runes := []rune(phrase)

	length := len(phrase)

	for index, letter := range phrase {
		runes[length-1-index] = letter
	}

	return string(runes)
}

func main() {
	fmt.Println(reversestring("Murder for a jar of red rum"))
}

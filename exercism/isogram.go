package main

import (
	"fmt"
	"strings"
)

func isogram(s string) bool {

	s = strings.ToLower(s)

	for i, letter := range s {
		duplicate := strings.ContainsRune(s[:i], letter)

		if duplicate == true {
			return false
		}
	}
	return true
}

func main() {

	word := "Uncopyrightable"

	word += " is"
	if isogram(word) {
		word += " not"
	}
	word += " a Isogram"

	fmt.Println(word)
}

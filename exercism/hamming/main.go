package main

import (
	"fmt"
)

func hamming(r string, s string) int {
	if len(s) != len(r) {
		fmt.Println(len(s))
		return -1
	}

	var t int
	for index := range s {
		if r[index] != s[index] {
			t++
		}
	}

	return t
}

func main() {

	strand_1 := "GAGCCTACTAACGGGAT"
	strand_2 := "CATCGTAATGACGGCCT"

	fmt.Println(strand_1)
	fmt.Println(strand_2)

	value := hamming(strand_1, strand_2)

	fmt.Println(value)
}

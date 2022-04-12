package main

import (
	"errors"
	"fmt"
)

func primefactors(num int) ([]int, error) {
	if num < 2 {
		return nil, errors.New("only natural numbers greater than 1")
	}

	var factors = []int{}

	for i := 2; i <= num; i++ {
		for num%i == 0 {
			factors = append(factors, i)
			num /= i
		}
	}

	return factors, nil
}

func main() {
	fmt.Println(primefactors(1))
	fmt.Println(primefactors(2))
	fmt.Println(primefactors(31))
	fmt.Println(primefactors(60))
	fmt.Println(primefactors(377))
}

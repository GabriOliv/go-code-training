package main

import (
	"fmt"
	"strings"
)

func luhn(creditnumber string) bool {
	// remove space
	creditnumber = strings.ReplaceAll(creditnumber, " ", "")

	// verify: length 1 or less
	if len(creditnumber) < 2 {
		return false
	}

	// verify: only number
	for _, letter := range creditnumber {
		if letter < 48 && letter > 57 {
			return false
		}
	}

	number := make([]int, len(creditnumber))

	// fill array
	for key, letter := range creditnumber {
		number[len(number)-1-key] = int(letter - 48)
	}

	// double every second number
	for key, digit := range number {
		if (key+1)%2 == 0 {
			number[key] = digit * 2

			if number[key] > 9 {
				number[key] -= 9
			}
		}
	}

	// sum of digits
	sum := 0
	for _, digit := range number {
		sum += digit
	}

	// verify: divisible by 10
	if sum%10 == 0 {
		return true
	}

	return false
}

func main() {

	fmt.Println(luhn("4539 3195 0343 6467"))

}

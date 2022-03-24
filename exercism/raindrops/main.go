package main

import (
	"fmt"
	"strconv"
)

func raindrops(number int) string {
	var rain string

	if number%3 == 0 {
		rain += "Pling"
	}
	if number%5 == 0 {
		rain += "Plang"
	}
	if number%7 == 0 {
		rain += "Plong"
	}
	if rain == "" {
		return strconv.Itoa(number)
	}

	return rain
}

func main() {

	var number int
	fmt.Scanln(&number)
	fmt.Println(raindrops(int(number)))

}

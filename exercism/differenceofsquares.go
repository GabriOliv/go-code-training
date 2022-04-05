package main

import "fmt"

func difference(number int) int {
	sum := 0
	sumofsquare := 0

	for number > 0 {
		sumofsquare += number * number
		sum += number
		number--
	}

	return (sum * sum) - sumofsquare
}

func main() {

	fmt.Println(difference(10))

}

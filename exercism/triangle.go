package main

import "fmt"

func category(a float64, b float64, c float64) string {
	if a == b && b == c {
		return "Equilateral"
	} else if a == b || a == c || b == c {
		return "Isosceles"
	} else {
		return "Scalene"
	}
}

func triangle(a float64, b float64, c float64) string {
	if a+b < c || a+c < b || b+c < a {
		return "not a Triangle"
	} else if a+b == c || a+c == b || b+c == a {
		return "Degenerate Triangle"
	}

	return category(a, b, c)
}

func main() {
	var sideA, sideB, sideC float64
	sideA = 6
	sideB = 12
	sideC = 5

	fmt.Println(sideA, sideB, sideC)
	fmt.Println(triangle(sideA, sideB, sideC))

}

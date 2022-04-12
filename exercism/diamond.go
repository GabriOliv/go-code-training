package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func draw(text string, size int, count int) {
	fmt.Print(strings.Repeat(text, size-count))
	fmt.Print(string(rune(count + 65)))

	if count != 0 {
		fmt.Print(strings.Repeat(text, (2*count - 1)))
		fmt.Print(string(rune(count + 65)))
	}

	fmt.Print(strings.Repeat(text, size-count))
	fmt.Println()
}

func diamond(letter string) error {
	// Match valid input
	if regexp.MustCompile(`^[A-Z]$`).MatchString(letter) == false {
		return errors.New("invalid input")
	}

	size := int(rune(letter[0])) - 65

	// Use Abs for count the negative part
	for i := size; int(math.Abs(float64(i))) <= size; i-- {
		draw(" ", size, size-int(math.Abs(float64(i))))
	}

	return nil
}

func main() {

	diamond("A")
	diamond("C")
	diamond("E")
}

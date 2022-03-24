package main

import (
	"fmt"
	"strings"

	"rsc.io/quote"
)

func main() {
	var message string
	message = "Hello, World!"
	name := "tester of this program"

	// Capitalize
	// name := strings.Title(name) // #deprecated
	name_array := strings.Split(name, " ")
	for key, value := range name_array {
		if len(value) <= 1 {
			value = strings.ToUpper(value)
		}
		value = strings.ToUpper(value[:1]) + value[1:]
		name_array[key] = value
	}
	name = strings.Join(name_array, " ")

	fmt.Println(fmt.Sprintf("\n%[1]s Hi, %[2]s\n", message, name))
	fmt.Println(quote.Glass())
	fmt.Println()
}

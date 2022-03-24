package main

import "fmt"

func main() {

	prefix := "One for "
	sufix := ", one for me."
	var name string

	fmt.Scanln(&name)

	if name == "" {
		name = "you"
	}

	fmt.Println(prefix + name + sufix)
}

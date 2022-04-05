package main

import (
	"bufio"
	"fmt"
	"os"
)

// Struct of Person
type programmer struct {
	name      string
	age       int
	languages []string
	contacts  []string
}

func main() {
	// new object
	person := new(programmer)

	// scan n print
	n, err := fmt.Scanln(&person.name)
	fmt.Println(n, err, person)

	n, err = fmt.Scanln(&person.age)
	fmt.Println(n, err, person)

	// error handling
	if err != nil {
		panic("no value for $USER")
	}

	// Declara buffer reader
	// default 4k size
	reader := bufio.NewReader(os.Stdin)

	line, isPrefix, err := reader.ReadLine()
	fmt.Println(line, isPrefix, err)

	fmt.Println(string(line))

}

package main

import (
	"errors"
	"fmt"
	"strings"
)

func acronym(phrase string) (string, error) {

	array := strings.Fields(strings.ReplaceAll(phrase, "-", " "))

	if len(array) < 2 {
		return "", errors.New("incorrect sentence")
	}

	for key, word := range array {
		array[key] = word[:1]
	}

	phrase = strings.ToUpper(strings.Join(array, ""))
	return phrase, nil
}

func main() {
	phrase, err := acronym("Three-letter  acronyms ")

	if err != nil {
		panic(err)
	}

	fmt.Println(phrase)
}

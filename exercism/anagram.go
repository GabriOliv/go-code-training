package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func anagram(word string, array []string) (string, error) {

	ord_word := strings.Split(word, "")
	sort.Strings(ord_word)

	for _, candidate := range array {

		ord_candidate := strings.Split(candidate, "")
		sort.Strings(ord_candidate)

		if strings.Join(ord_word, "") == strings.Join(ord_candidate, "") {
			return candidate, nil
		}
	}

	return "", errors.New("now found")
}

func main() {
	word := "listen"
	candidates := []string{"enlists", "google", "inlets", "banana"}

	word, err := anagram(word, candidates)

	if err != nil {
		panic(err)
	}

	fmt.Println(word)
}

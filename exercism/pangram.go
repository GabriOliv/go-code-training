package main

import (
	"errors"
	"fmt"
	"strings"
)

func pangram(phrase string) (bool, error) {

	phrase = strings.ToLower(strings.ReplaceAll(phrase, " ", ""))
	// summation of alphabet runes
	runesummation := 2847

	for key, letter := range phrase {

		// verify non-letter characters
		if letter < 97 && letter > 122 {
			return false, errors.New("invalid character")
		}

		// verify alphabet runes
		if strings.ContainsRune(phrase, rune(key+97)) == true {
			runesummation -= (key + 97)
		}
	}

	if runesummation == 0 {
		return true, nil
	}

	return false, nil
}

func main() {

	ispangram, err := pangram("The quick brown fox jumps over the lazy dog")

	if err != nil {
		panic(err)
	}

	fmt.Println(ispangram)
}

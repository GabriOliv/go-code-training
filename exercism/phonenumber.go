package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func phoneNumber(input string) (string, error) {
	// return 6139950253

	// match Possible inputs
	match := regexp.MustCompile(`^(\+1\s|1\s)*((\([2-9]{1}[0-9]{2}\)|[2-9]{1}[0-9]{2})\s([2-9]{1}[0-9]{2})\s([0-9]{4})|(\([2-9]{1}[0-9]{2}\)|[2-9]{1}[0-9]{2})\-([2-9]{1}[0-9]{2})\-([0-9]{4})|(\([2-9]{1}[0-9]{2}\)|[2-9]{1}[0-9]{2})\.([2-9]{1}[0-9]{2})\.([0-9]{4}))$`)

	// Not Match Error
	if match.MatchString(input) == false {
		return "", errors.New("invalid NANP telephone number")
	}

	input = strings.ReplaceAll(input, "-", " ")
	input = strings.ReplaceAll(input, ".", " ")
	input = strings.ReplaceAll(input, "(", " ")
	input = strings.ReplaceAll(input, ")", " ")

	// Submatch regex
	match = regexp.MustCompile(`(?P<part1>[2-9]{1}[0-9]{2})\s*(?P<part2>[2-9]{1}[0-9]{2})\s*(?P<part3>[0-9]{4})`)

	// Join submatch
	input = strings.Join(match.FindStringSubmatch(input)[1:], "")

	return input, nil
}

func areaCode(input string) (string, error) {
	// return 613

	input, err := phoneNumber(input)
	if err != nil {
		return "", err
	}

	return input[:3], nil
}

func format(input string) (string, error) {
	// (613) 995-0253

	input, err := phoneNumber(input)
	if err != nil {
		return "", err
	}

	return ("(" + input[:3] + ") " + input[3:6] + "-" + input[6:]), nil
}

func errorHandling(err error) {
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}
}

func main() {

	inputs := []string{
		// true
		"+1 (613)-995-0253",
		"613-995-0253",
		"1 613 995 0253",
		"613.995.0253",
		// false
		// "+2 (613)-995-0253",
		// "013-995-0253",
		// "1 613 995-0253",
		// "613.9950253",
	}

	for _, elem := range inputs {

		fmt.Println()
		fmt.Println("Source:", elem)

		output, err := phoneNumber(elem)
		errorHandling(err)
		fmt.Println("\tRaw:\t", output)

		output, err = format(elem)
		errorHandling(err)
		fmt.Println("\tFormat:\t", output)

		output, err = areaCode(elem)
		errorHandling(err)
		fmt.Println("\tArea:\t", output)

	}

	fmt.Println()
}

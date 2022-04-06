package main

import (
	"errors"
	"fmt"
	"strings"
)

func rnatranscription(dna string) ([]int, error) {

	dna = strings.ToUpper(dna)

	count := make([]int, 4)

	for _, nucleotide := range dna {
		switch {
		case nucleotide == 65:
			count[0]++
		case nucleotide == 67:
			count[1]++
		case nucleotide == 71:
			count[2]++
		case nucleotide == 84:
			count[3]++
		default:
			return make([]int, 4), errors.New("wrong sequence of nucleotides")
		}
	}

	return count, nil
}

func main() {

	label := `
DNA:
	(A) adenine,
	(C) cytosine 
	(G) guanine 
	(T) thymine 
`
	count_label := []string{"A", "C", "G", "T"}

	fmt.Println(label)

	dna := "gattaca"

	count, err := rnatranscription(dna)

	if err != nil {
		panic(err)
	}

	for index, nucleotide := range count {
		fmt.Print(count_label[index], ":", nucleotide, " ")
	}
	fmt.Println()

}

package main

import (
	"errors"
	"fmt"
	"strings"
)

func rnatranscription(dna string) (string, error) {

	dna = strings.ToUpper(dna)

	rna := ""

	for _, nucleotide := range dna {
		switch {
		case nucleotide == 71:
			rna += string(rune(67))
		case nucleotide == 67:
			rna += string(rune(71))
		case nucleotide == 84:
			rna += string(rune(65))
		case nucleotide == 65:
			rna += string(rune(85))
		default:
			return "", errors.New("wrong sequence of nucleotides")
		}
	}

	return rna, nil
}

func main() {

	label := `
DNA:
	(A) adenine,
	(C) cytosine 
	(G) guanine 
	(T) thymine 
RNA:
	(A) adenine 
	(C) cytosine
	(G) guanine
	(U) uracil

Transcription:
	G -> C
	C -> G
	T -> A
	A -> U
`

	fmt.Println(label)

	dna := "gctagcttagc"

	rna, err := rnatranscription(dna)

	if err != nil {
		panic(err)
	}

	fmt.Println("DNA:", dna)
	fmt.Println()
	fmt.Println("RNA:", rna)
}

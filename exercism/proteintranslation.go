package main

import (
	"errors"
	"fmt"
	"regexp"
)

func proteinTranslation(rna string) ([]string, error) {

	// Match valid rna sequence
	if regexp.MustCompile(`^((A|C|G|U){3})+$`).MatchString(rna) == false {
		return nil, errors.New("invalid rna sequence")
	}

	// Split rna
	var proteins = []string{}
	for key := range rna {
		if (key+1)%3 == 0 {
			proteins = append(proteins, rna[(key-2):key+1])
		}
	}

	// Translate sequence
	for key, sequence := range proteins {
		switch {
		case sequence == "AUG":
			proteins[key] = "Methionine"
		case sequence == "UUU":
			proteins[key] = "Phenylalanine"
		case sequence == "UUC":
			proteins[key] = "Phenylalanine"
		case sequence == "UUA":
			proteins[key] = "Leucine"
		case sequence == "UUG":
			proteins[key] = "Leucine"
		case sequence == "UCU":
			proteins[key] = "Serine"
		case sequence == "UCC":
			proteins[key] = "Serine"
		case sequence == "UCA":
			proteins[key] = "Serine"
		case sequence == "UCG":
			proteins[key] = "Serine"
		case sequence == "UAU":
			proteins[key] = "Tyrosine"
		case sequence == "UAC":
			proteins[key] = "Tyrosine"
		case sequence == "UGU":
			proteins[key] = "Cysteine"
		case sequence == "UGC":
			proteins[key] = "Cysteine"
		case sequence == "UGG":
			proteins[key] = "Tryptophan"
		// STOP Protein
		case sequence == "UAA":
			return proteins[:key], nil
		case sequence == "UAG":
			return proteins[:key], nil
		case sequence == "UGA":
			return proteins[:key], nil
		// Unindexed Protein
		default:
			proteins[key] = "Unindexed"
		}
	}

	return proteins, nil
}

func main() {

	rna := []string{
		"AUGUUUUCU",
		"AUGUUUUCUUAAAUG",
		"AUGUUUUCUAUGUAA",
		"UCUUAAAUGUUUUCUUAAAUG",
		"UCUUAAAUGUUUUCUUAAAUGUCU",
		"AUGUUUUUUUCUUAAAUG",
	}

	for _, strand := range rna {

		sequence, err := proteinTranslation(strand)
		if err != nil {
			panic(err)
		}

		fmt.Println(strand, sequence)
	}
}

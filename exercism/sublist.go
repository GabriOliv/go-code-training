package main

import "fmt"

func equal(list1 []int, list2 []int) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}

	}
	return true
}

func sublist(list1 []int, list2 []int) int {
	// -1 A<B  sublist
	//  0 A=B  equal
	//  1 A>B  superlist
	//  2 A!=B different

	len1 := len(list1)
	len2 := len(list2)

	if len1 < len2 {
		for i := 0; i <= (len2 - len1); i++ {
			if equal(list1, list2[i:(len1+i)]) == true {
				// -1 A<B  sublist
				return -1
			}
		}
	} else if len1 > len2 {
		for i := 0; i <= (len1 - len2); i++ {
			if equal(list2, list1[i:(len2+i)]) == true {
				//  1 A>B  superlist
				return 1
			}
		}
	} else if len1 == len2 {
		if equal(list1, list2) == true {
			//  0 A=B  equal
			return 0
		}
	}

	//  2 A!=B different
	return 2
}

func main() {
	table := [][][]int{
		// Test 1
		{{1, 2, 3}, {1, 2, 3, 4, 5}},
		// Test 2
		{{3, 4, 5}, {1, 2, 3, 4, 5}},
		// Test 3
		{{3, 4}, {1, 2, 3, 4, 5}},
		// Test 4
		{{1, 2, 3}, {1, 2, 3}},
		// Test 5
		{{1, 2, 3, 4, 5}, {2, 3, 4}},
		// Test 6
		{{1, 2, 4}, {1, 2, 3, 4, 5}},
	}

	for _, test := range table {
		fmt.Print("A: ", test[0], "\t\tB: ", test[1])

		test := sublist(test[0], test[1])

		switch {
		case test == -1:
			fmt.Println("\t\tA is a sublist of B")
		case test == 0:
			fmt.Println("\t\tA is a equal of B")
		case test == 1:
			fmt.Println("\t\tA is a superlist of B")
		default:
			fmt.Println("\t\tA is not a superlist of, sublist of or equal to B")
		}
	}

}

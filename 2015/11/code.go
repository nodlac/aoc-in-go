package main

import (
	"fmt"
	"strings"
)

// Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
// Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
// Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.

func isValid(input []byte, originalString string) bool {
	if string(input) == originalString {
		return false
	}

	evilLetters := "iol"
	if strings.ContainsAny(string(input), evilLetters) {
		return false
	}

	hasRun := false
	for i := 0; i < len(input)-3; i++ {
		fmt.Printf("%d, %d, %d\n", input[i], input[i+1], input[i+2])
		if input[i]+1 == input[i+1] && input[i]+2 == input[i+2] {
			hasRun = true
		}
	}
	if !hasRun {
		return false
	}

	hasDoubled := true
	for i := 0; i < len(input)-1; i++ {
	}
	if !hasDoubled {
		return false
	}

	return true
}

func main() {
	// z = 122
	// a = 97
	// input := "cqjxjnds"
	input := "cqjxjnds"
	converted := []byte(input)

	for !isValid(converted, input) {
		fmt.Println(converted)
		for i := len(converted) - 1; i >= 0; i-- {
			converted[i] += 1
			if converted[i] > 122 {
				converted[i] = 97
				continue
			}
			break
		}
	}

	input = string(converted)

	fmt.Println(input)
}

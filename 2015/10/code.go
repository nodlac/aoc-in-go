package main

import (
	"fmt"
	"strings"
	"strconv"
)

func getIteration(input string) string {
	var builder strings.Builder
	for i := 0; i < len(input); {
		current := input[i]
		countCurr := 1
		for i = i+1; i < len(input); i++ {
			if input[i] != current {
				break
			}
			countCurr += 1
		}
		builder.WriteString(strconv.Itoa(countCurr))
		builder.WriteByte(current)
	}

	return builder.String()
}

func main() {
	
	// input := "1"
	input := "3113322113"
	for i := 0; i < 50; i++ {
		input = getIteration(input)
	}
	fmt.Println(len(input))
}

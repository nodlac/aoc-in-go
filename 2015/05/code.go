package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
	good := 0
	// file, err := os.Open("./input-example.txt")
	// file, err := os.Open("./input-example2.txt")
	file, err := os.Open("./input-user.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error opening file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		hasPair := false
		hasRepeating := false
		// fmt.Println(line)
	matchSearch:
		for i := 0; i < len(line)-3; i++ {
			pair := line[i : i+2]
			for j := i + 2; j < len(line)-1; j++ {
				potentialMatch := line[j : j+2]
				// fmt.Printf("pair = %v, potential = %v\n", pair, potentialMatch)
				if pair == potentialMatch {
					hasPair = true
					break matchSearch
				}
			}
		}
		if !hasPair {
			continue
		}
		for i := 0; i < len(line)-2; i++ {
			// fmt.Printf("pair = %v, potential = %v\n", line[i], line[i+2])
			if line[i] == line[i+2] {
				hasRepeating = true
				break
			}
		}
		if !hasRepeating {
			continue
		}
		good += 1
	}

	// vowels := "aeiouAEIOU"
	// line:
	// for scanner.Scan() {
	// 	vs := make(map[rune]int)
	// 	vCount := 0
	// 	containsDouble := false
	//
	// 	line := scanner.Text()
	// 	runes := []rune(line)
	// 	for i := 0; i < len(runes); i++ {
	// 		if i < len(runes) - 1 {
	// 			if !containsDouble {
	// 				containsDouble = runes[i] == runes[i+1]
	// 			}
	//
	// 			proxDeux := string(runes[i:i+2])
	//
	// 			if proxDeux == "ab"|| proxDeux ==  "cd" || proxDeux == "pq" || proxDeux == "xy" {
	// 				fmt.Println("failed proxDuex")
	// 				continue line
	// 			}
	// 		}
	// 		if strings.ContainsRune(vowels, runes[i]) {
	// 			vCount += 1
	// 			vs[runes[i]] += 1
	// 		}
	// 	}
	// 	if vCount < 3    {
	// 		fmt.Println("failed 3 Vowel")
	// 		continue
	// 	}
	// 	if !containsDouble {
	// 		fmt.Println("failed containsDouble")
	// 		continue
	// 	}
	// 	good += 1
	// 	fmt.Println(good)
	// }
	fmt.Printf("Total good lines = %d\n", good)

}

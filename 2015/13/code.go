package main

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func getPermutations(options []string) [][]string {
	var result [][]string
	backtrace(options, 0, &result)
	return result
}

func backtrace(options []string, start int, result *[][]string) {
	if start == len(options) {
		temp := make([]string, len(options))
		copy(temp, options)
		*result = append(*result, temp)
		return
	}
	for i := start; i < len(options); i++ {
		options[start], options[i] = options[i], options[start]
		backtrace(options, start+1, result)
		options[start], options[i] = options[i], options[start]

	}

}

type person struct {
	name   string
	people map[string]int
}

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user inputA
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	happinessMap := make(map[string]person)

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, " ")
		happiness, err := strconv.Atoi(parts[3])
		if err != nil {
			panic(err)
		}
		if parts[2] == "lose" {
			happiness = happiness * -1
		}

		name := parts[0]
		partner := parts[len(parts)-1]
		partner = partner[:len(partner)-1]

		h := happinessMap[name]
		if h.name == "" {
			h.name = name
		}
		if h.people == nil {
			h.people = make(map[string]int)
		}
		h.people[partner] = happiness
		happinessMap[name] = h

	}

	happinessMap["me"] = person{name: "me"}

	keys := make([]string, 0, len(happinessMap))

	for k, _ := range happinessMap {
		keys = append(keys, k)
	}


	permuts := getPermutations(keys)

	highest := 0
	for _, p := range permuts {
		sum := 0

		for i := 0; i < len(p); i++ {
			name := p[i]
			var left string
			var right string
			if i-1 < 0 {
				left = p[len(p)-1]
			} else {
				left = p[i-1]
			}
			if i+1 > len(p)-1 {
				right = p[0]
			} else {
				right = p[i+1]
			}
			sum += happinessMap[name].people[left]
			sum += happinessMap[name].people[right]
		}

		if sum > highest {
			highest = sum
		}
	}

	// solve part 1 here
	return highest
}

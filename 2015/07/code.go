package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout

type wire struct {
	inputs []string
	gate   string
	value  int
	shift  int
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	// file, err := os.Open("./input-user.txt")
	file, err := os.Open("./input-user.txt") // answer 492
	if err != nil {
		fmt.Println("Error loading file")
	}
	scanner := bufio.NewScanner(file)

	circutMap := make(map[string]wire)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " -> ")

		input := parts[0]
		target := parts[1]

		inputParts := strings.Split(input, " ")

		switch {
		case len(inputParts) == 1:
			val, eVal := strconv.Atoi(input)
			if eVal != nil {
				circutMap[target] = wire{inputs: []string{inputParts[0]}}
			} else {
				circutMap[target] = wire{value: val}
			}

		case inputParts[1] == "AND":
			circutMap[target] = wire{
				gate:   inputParts[1],
				inputs: []string{inputParts[0], inputParts[2]},
			}
		case inputParts[1] == "OR":
			circutMap[target] = wire{
				gate:   inputParts[1],
				inputs: []string{inputParts[0], inputParts[2]},
			}

		case inputParts[1] == "LSHIFT":
			s, sErr := strconv.Atoi(inputParts[2])
			if sErr != nil {
				fmt.Println("error converting shift")
			}
			circutMap[target] = wire{
				gate:   inputParts[1],
				shift:  s,
				inputs: []string{inputParts[0]},
			}
		case inputParts[1] == "RSHIFT":
			s, sErr := strconv.Atoi(inputParts[2])
			if sErr != nil {
				fmt.Println("error converting shift")
			}
			circutMap[target] = wire{
				gate:   inputParts[1],
				shift:  s,
				inputs: []string{inputParts[0]},
			}
		case inputParts[0] == "NOT":
			circutMap[target] = wire{
				gate:   inputParts[0],
				inputs: []string{inputParts[1]},
			}
		default:
			fmt.Printf("Uncovered case %v\n", line)
			return -1
		}
	}

	memo := make(map[string]int)
	r := getValue("a", circutMap, memo)
	memo = make(map[string]int)
	memo["b"] = r
	return getValue("a", circutMap, memo)
}

func getValue(name string, circutMap map[string]wire, memo map[string]int) int {
	if val, ok := memo[name]; ok {
		return val
	}

	if val, err := strconv.Atoi(name); err == nil {
		return val
	}

	w := circutMap[name]

	if w.value != 0 && w.gate == "" && len(w.inputs) == 0 {
		memo[name] = w.value
		return w.value
	}
	var result int
	switch w.gate {
	case "":
		if len(w.inputs) > 0 {
		result = getValue(w.inputs[0], circutMap, memo)
		} else {
			result = w.value
		}
	case "AND":
		result = getValue(w.inputs[0], circutMap, memo) & getValue(w.inputs[1], circutMap, memo)

	case "OR":
		result = getValue(w.inputs[0], circutMap, memo) | getValue(w.inputs[1], circutMap, memo)

	case "LSHIFT":
		result = getValue(w.inputs[0], circutMap, memo) << w.shift

	case "RSHIFT":
		result = getValue(w.inputs[0], circutMap, memo) >> w.shift

	case "NOT":
		result = int(^uint16(getValue(w.inputs[0], circutMap, memo)))
	default:
		fmt.Printf("Uncovered case %v\n", w)
		return -1
	}
	result &= 0xFFFF
	memo[name] = result 
	return result
}

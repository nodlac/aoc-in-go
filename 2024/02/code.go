package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
	"strconv"
	"bufio"
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


func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	safe := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		totalDiff := 0
		for i := 0; i < len(parts) - 1; i++ {
			j := i+1
			i_val, err_i := strconv.Atoi(parts[i]) 
			j_val, err_j := strconv.Atoi(parts[j]) 
			if err_i != nil || err_j != nil {
				break
			}
			diff := i_val - j_val
			if diff == 0 {
				break
			}
			if diff > 3 || diff < -3 {
				break 
			}
			if totalDiff < 0 && diff > 0 {
				break
			}
			if totalDiff > 0 && diff < 0 {
				break
			}
			if j == len(parts)-1 {
				safe++ 
			}
			totalDiff += diff
		}
	}
	return safe
}

package main

import (
	"sort"
	"bufio"
	"fmt"
	"regexp"
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
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	list1 := make([]int, 0, 100)
	list2 := make([]int, 0, 100)
	scanner := bufio.NewScanner(strings.NewReader(input))

	re := regexp.MustCompile(`\s+`)

	for scanner.Scan() {
		line := scanner.Text()
		parts := re.Split(line, -1)
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		list1 = append(list1, a)
		list2 = append(list2, b)
	}
	sort.Ints(list1)
	sort.Ints(list2)

	if part2 {
		likeness := 0
		freq := make(map[int]int)
			for _, v := range list2 {
			freq[v]++
		}
		for _, v := range list1 {
			likeness += freq[v] * v
		}
		return strconv.Itoa(likeness)
	}

	sum := 0
	for i, v := range list1 {
		diff := v - list2[i]
		if diff < 0 {
		diff = -diff
	}
		sum += diff
	}

	fmt.Println(sum)
	return strconv.Itoa(sum)
}

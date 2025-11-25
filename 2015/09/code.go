package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

type node struct {
	name string
	edges map[string]int
}

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
	scanner := bufio.NewScanner(strings.NewReader(input))

	nodeMap := make(map[string]node)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		expressions := strings.Split(line, " = ")
		dist, err := strconv.Atoi(expressions[1])
		if err != nil {
			fmt.Println("Error converting distance")
			return -1
		}
		locations := strings.Split(expressions[0], " to ")

		if nodeMap[locations[0]].name == "" {
			nodeMap[locations[0]] = node{name: locations[0], edges:make(map[string]int)}
		}
		curNode := nodeMap[locations[0]]
		curNode.edges[locations[1]] = dist

		if nodeMap[locations[1]].name == "" {
			nodeMap[locations[1]] = node{name: locations[1], edges:make(map[string]int)}
		}
		curNode = nodeMap[locations[1]]
		curNode.edges[locations[0]] = dist
	}


	var keys []string
	for key := range(nodeMap){
		keys = append(keys, key)
	}

	fmt.Println(keys)
	// solve part 1 here
	return 42
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Light struct {
	x int
	y int
}

func main() {
	grid := make(map[Light]int)
	// file, err := os.Open("input-example.txt")
	file, err := os.Open("input-user.txt")
	if err != nil {
		fmt.Println("error reading file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		reOp := regexp.MustCompile("^[^0-9]*")
		operation := reOp.FindStringSubmatch(line)
		if operation == nil {
			fmt.Println("no match")
		}
		reCoords := regexp.MustCompile("[0-9]+")
		coordinates := reCoords.FindAllStringSubmatch(line, -1)
		if coordinates == nil {
			fmt.Println("no match")
		}
		if len(coordinates) < 4 {
			fmt.Println("not enough coords")
			return
		}
		if len(operation) == 0 {
			fmt.Println("error no match found")
			fmt.Println(line)
			continue
		}
		xStart, errXS := strconv.Atoi(coordinates[0][0])
		xEnd, errXE := strconv.Atoi(coordinates[2][0])
		yStart, errYS := strconv.Atoi(coordinates[1][0])
		yEnd, errYE := strconv.Atoi(coordinates[3][0])

		if errXS != nil || errXE != nil || errYS != nil || errYE != nil {
			fmt.Println("error with conversions")
			return
		}

		for i := xStart; i <= xEnd; i++ {
			for j := yStart; j <= yEnd; j++ {
				// fmt.Printf("%d,%d\n", i, j)

				current := Light{x: i, y: j}
				switch operation[0] {
				case "toggle ":
					grid[current] += 2
				case "turn off ":
					grid[current] -= 1
					if grid[current] < 0 {
						grid[current] = 0
					}
				case "turn on ":
					grid[current] += 1
				}
			}
		}
	}
	sum := 0
	for _, v := range grid {
		sum += v

	}
	fmt.Printf("total on = %d\n", sum)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	paperTotal := 0
	ribbonTotal := 0

	// file, err := os.Open("./example_input.txt")
	file, err := os.Open("./user_inputx.txt")
	if err != nil {
		fmt.Println("Couldn't read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "x")
		if len(parts) < 3 {
			fmt.Println("too few dimensions")
			return
		}
		l, lErr := strconv.Atoi(parts[0])
		w, wErr := strconv.Atoi(parts[1])
		h, hErr := strconv.Atoi(parts[2])

		if lErr != nil || wErr != nil || hErr != nil {
			fmt.Println("dimension error")
			return
		}
		m, c := leastTwo(l, w, h)
		additional := 2*l*w + 2*w*h + 2*h*l + m*c
		paperTotal += additional

		ribbon := m*2 + c*2 + l*w*h
		fmt.Println(ribbon)
		ribbonTotal += ribbon
	}

	fmt.Printf("Total Paper %d\n", paperTotal)
	fmt.Printf("Total Ribbon %d\n", ribbonTotal)

}

func leastTwo(a, b, c int) (int, int) {
	// Start by assuming a < b < c
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	return a, b
}

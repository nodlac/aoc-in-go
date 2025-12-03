package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sue struct {
	id          string
	score       int
	children    *int
	cats        *int
	samoyeds    *int
	pomeranians *int
	akitas      *int
	vizslas     *int
	goldfish    *int
	trees       *int
	cars        *int
	perfumes    *int
}

func intPtr(i int) *int {
	return &i
}

func main() {

	gifter := sue{
		id:          "0",
		children:    intPtr(3),
		cats:        intPtr(7),
		samoyeds:    intPtr(2),
		pomeranians: intPtr(3),
		akitas:      intPtr(0),
		vizslas:     intPtr(0),
		goldfish:    intPtr(5),
		trees:       intPtr(3),
		cars:        intPtr(2),
		perfumes:    intPtr(1),
	}

	filename := "./input-user.txt"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	matchScore := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ":", "")
		line = strings.ReplaceAll(line, ",", "")
		parts := strings.Split(line, " ")
		name := parts[1]
		temp := sue{id: string(name)}
		for i := 2; i < len(parts)-1; i += 2 {
			val, err := strconv.Atoi(parts[i+1])
			if err != nil {
				panic(err)
			}
			switch parts[i] {
			case "children":
				temp.children = intPtr(val)
			case "samoyeds":
				temp.samoyeds = intPtr(val)
			case "cats":
				temp.cats = intPtr(val)
			case "pomeranians":
				temp.pomeranians = intPtr(val)
			case "akitas":
				temp.akitas = intPtr(val)
			case "vizslas":
				temp.vizslas = intPtr(val)
			case "goldfish":
				temp.goldfish = intPtr(val)
			case "trees":
				temp.trees = intPtr(val)
			case "cars":
				temp.cars = intPtr(val)
			case "perfumes":
				temp.perfumes = intPtr(val)
			}
		}

		score := 0

		if temp.children != nil {
			if *temp.children != *gifter.children {
				continue
			} else {
				score += 1
			}
		}
		if temp.cats != nil {
			if *temp.cats <= *gifter.cats {
				continue
			} else {
				score += 1
			}
		}
		if temp.samoyeds != nil {
			if *temp.samoyeds != *gifter.samoyeds {
				continue
			} else {
				score += 1
			}
		}
		if temp.pomeranians != nil {
			if *temp.pomeranians >= *gifter.pomeranians {
				continue
			} else {
				score += 1
			}
		}
		if temp.vizslas != nil {
			if *temp.vizslas != *gifter.vizslas {
				continue
			} else {
				score += 1
			}
		}
		if temp.goldfish != nil {
			if *temp.goldfish >= *gifter.goldfish {
				continue
			} else {
				score += 1
			}
		}
		if temp.akitas != nil {
			if *temp.akitas != *gifter.akitas {
				continue
			} else {
				score += 1
			}
		}
		if temp.trees != nil {
			if *temp.trees <= *gifter.trees {
				continue
			} else {
				score += 1
			}
		}
		if temp.cars != nil {
			if *temp.cars != *gifter.cars {
				continue
			} else {
				score += 1
			}
		}
		if temp.perfumes != nil {
			if *temp.perfumes != *gifter.perfumes {
				continue
			} else {
				score += 1
			}
		}

		matchScore[temp.id] = score
	}

	fmt.Println(matchScore)

}

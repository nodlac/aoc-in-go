package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coords struct {
	X, Y int
}

func (c *Coords) AddY(amount int) {
	c.Y += amount
}

func (c *Coords) AddX(amount int) {
	c.X += amount
}

func main() {
	// file, err := os.Open("./input-example.txt")
	file, err := os.Open("./input-user.txt")
	if err != nil {
		fmt.Println("error loading file")
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		santa := Coords{X: 0, Y: 0}
		robot := Coords{X: 0, Y: 0}
		houses := make(map[Coords]int)
		houses[Coords{0, 0}] += 1
		text := scanner.Text()
		for i := 0; i < len(text); i++ {
			if i%2 == 0 {
				switch text[i] {
				case '<':
					robot.AddX(-1)
				case '>':
					robot.AddX(1)
				case '^':
					robot.AddY(1)
				case 'v':
					robot.AddY(-1)
				}
				houses[robot] += 1
			} else {
				switch text[i] {
				case '<':
					santa.AddX(-1)
				case '>':
					santa.AddX(1)
				case '^':
					santa.AddY(1)
				case 'v':
					santa.AddY(-1)
				}
				houses[santa] += 1
			}
		}
		fmt.Println(len(houses))
	}

}

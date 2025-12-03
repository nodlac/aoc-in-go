package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func getPermutations(items int, volume int) [][]int {
	numer := 1
	demur := 1
	for i := items + volume - 1; i > volume; i-- {
		numer *= i
	}
	for i := items - 1; i > 0; i-- {
		demur *= i
	}
	totalPermuts := numer / demur
	result := make([][]int, 0, totalPermuts)
	composition := make([]int, items)
	finished := false
	for !finished {
		sum := 0
		for i := 0; i < len(composition); i++ {
			sum += composition[i]
		}
		if sum == 100 {
			temp := make([]int, len(composition))
			copy(temp, composition)
			result = append(result, temp)
		}

		finished = helper(0, volume, composition)
	}
	return result

}

func helper(i int, volume int, composition []int) bool {
	composition[i] += 1
	if composition[len(composition)-1] == 100 {
		return true
	}
	if composition[i] == volume {
		composition[i] = 0
		helper(i+1, volume, composition)
	}
	return false
}

func main() {

	filename := "./input-user.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	ingredients := []ingredient{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		parts[1] = strings.ReplaceAll(parts[1], ",", "")
		qualities := strings.Split(parts[1], " ")
		capacity, err := strconv.Atoi(qualities[1])
		if err != nil {
			panic(err)
		}

		durability, err := strconv.Atoi(qualities[3])
		if err != nil {
			panic(err)
		}

		flavor, err := strconv.Atoi(qualities[5])
		if err != nil {
			panic(err)
		}

		texture, err := strconv.Atoi(qualities[7])
		if err != nil {
			panic(err)
		}

		calories, err := strconv.Atoi(qualities[9])
		if err != nil {
			panic(err)
		}

		temp := ingredient{parts[0], capacity, durability, flavor, texture, calories}

		ingredients = append(ingredients, temp)

	}

	permuts := getPermutations(len(ingredients), 100)

	highest := 0
	for _, p := range permuts {
		score := 0
		sumC := 0
		sumD := 0
		sumF := 0
		sumT := 0
		sumCal := 0

		for j := 0; j < len(ingredients); j++ {
			curr := ingredients[j]
			sumC += curr.capacity * p[j]

			sumD += curr.durability * p[j]

			sumF += curr.flavor * p[j]

			sumT += curr.texture * p[j]
			
			sumCal += curr.calories * p[j]

		}
		sumC = int(math.Max(float64(sumC), float64(0)))
		sumD = int(math.Max(float64(sumD), float64(0)))
		sumF = int(math.Max(float64(sumF), float64(0)))
		sumT = int(math.Max(float64(sumT), float64(0)))
		fmt.Println(p)
		fmt.Println(score)
		sum := sumC * sumD * sumF * sumT
		score += sum
		if score > highest && sumCal == 500 {
			highest = score
		}
	}

	fmt.Println(highest)

}

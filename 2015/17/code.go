package main

import (
	"fmt"
)

func getCombos(containers []int, combos *[][]int, slots int)  {
	var current []int

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(current) == slots {
			temp := make([]int, slots)
			copy(temp, current)
			*combos = append(*combos, temp)
			return
		}

		for i := start; i < len(containers); i++ {
			current = append(current, containers[i])
			backtrack(i + 1) // move forward only
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
}

// func getPermuts(containers []int, n int, combos *[][]int, positions int) {
// 	for i := n; i < len(containers)-1; i++ {
// 		for j := i + 1; j < len(containers); j++ {
// 			containers[i], containers[j] = containers[j], containers[i]
// 			temp := make([]int, positions)
// 			copy(temp, containers)
// 			*combos = append(*combos, temp)
// 			if positions > 1 {
// 				getPermuts(containers, n+1, combos, positions)
// 			}
// 			containers[i], containers[j] = containers[j], containers[i]
// 		}
// 	}
// }

// func getValidPermuts(combos [][]int, volume int) [][]int {
// 	// comboCopy := make([][]int, len(combos))
// 	// copy(comboCopy, combos)
// 	for i, p := range combos {
// 		fmt.Println(p)
// 		sum := 0
// 		for _, v := range p {
// 			sum += v
// 			if v != volume {
// 				combos[i] == 0
// 			}
// 		}
// 	}
// 	return combos
// }

func main() {
	containerSizes := []int{
		// 20, 15, 10, 5, 5,
		// 1, 2, 3,

		43,
		3,
		4,
		10,
		21,
		44,
		4,
		6,
		47,
		41,
		34,
		17,
		17,
		44,
		36,
		31,
		46,
		9,
		27,
		38,

	}

	totalVariants := 0
	for k := len(containerSizes); k > 0; k-- {
		iterNum := 1
		iterDenum := 1
		for n := len(containerSizes); n > k; n-- {
			iterNum *= n
		}
		for d := len(containerSizes)-k; d > 0; d-- {
			iterDenum *= d
		}

		totalVariants += iterNum / iterDenum
	}

		
	combos := make([][]int, 0, totalVariants)

	for l := len(containerSizes); l > 0; l-- {
		getCombos(containerSizes, &combos, l)
	}


	// volume := 25
	volume := 150

	validPermuts := 0
	for _, v := range combos {
		sum := 0
		for _, c := range v {
			sum += c
		}
		if sum == volume {
			fmt.Println(v)
			validPermuts += 1
		}
	}

	fmt.Println(combos)
	fmt.Println("len combos ", len(combos))
	fmt.Println("total vars ", totalVariants)
	fmt.Println("len validPermuts ", validPermuts)
}

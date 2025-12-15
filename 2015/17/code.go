package main

import (
	"fmt"
)

func getPermuts(containers []int, n int, combos *[][]int, positions int) {
	for i := n; i < len(containers)-1; i++ {
		for j := i + 1; j < len(containers); j++ {
			containers[i], containers[j] = containers[j], containers[i]
			temp := make([]int, positions)
			copy(temp, containers)
			*combos = append(*combos, temp)
			if positions > 1 {
				getPermuts(containers, n+1, combos, positions)
			}
			containers[i], containers[j] = containers[j], containers[i]
		}
	}
}

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
		1, 2, 3,

		// real sizes
		// 43,
		// 3,
		// 4,
		// 10,
		// 21,
		// 44,
		// 4,
		// 6,
		// 47,
		// 41,
		// 34,
		// 17,
		// 17,
		// 44,
		// 36,
		// 31,
		// 46,
		// 9,
		// 27,
		// 38,
	}

	totalVariants := 0
	for i := len(containerSizes); i > 0; i-- {
		iterVars := 1
		for j := len(containerSizes); j > (len(containerSizes) - i); j-- {
			iterVars *= j
		}
		iterDenom
		totalVariants += iterVars
	}

	combos := make([][]int, 0, totalVariants)

	for l := len(containerSizes); l > 0; l-- {
		getPermuts(containerSizes, 0, &combos, l)
	}

	fmt.Println(combos)

	volume := 5
	// volume := 150

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

	fmt.Println(totalVariants)
	fmt.Println(validPermuts)
}

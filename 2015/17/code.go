package main

import (
	"fmt"
)

func getCombos(containers []int, combos *[][]int, listLength int) {
	for l := listLength; l > 0; l-- {
		for i := 0; i < l-1; i++ {
			for j := i; j < l; j++ {
				containers[i], containers[j] = containers[j], containers[i]
				temp := make([]int, 0, l)
				temp = copy(temp, containers)
				*combos = append(*combos, temp)
				containers[i], containers[j] = containers[j], containers[i]
			}
		}
	}
}

func main() {
	containerSizes := []int{
		20, 15, 10, 5, 5,

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

	totalVariants := 1
	for i := len(containerSizes); i > 1; i-- {
		totalVariants *= i
	}

	combos := make([][]int, 0, totalVariants)

	getCombos(containerSizes, &combos, len(containerSizes))

	fmt.Println(combos)
	//
	// volume := 25
	// volume := 150

}

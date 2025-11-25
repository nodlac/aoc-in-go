package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, eFile := os.Open("./input-user.txt")
	if eFile != nil {
		fmt.Println("error opening file")
		return
	}

	scanner := bufio.NewScanner(file)

	sumTotal := 0
	sumText := 0

	// // decode
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	if len(line) == 0 {
	// 		continue
	// 	}cont
	// 	sumTotal += len(line)
	// 	formatted, err := strconv.Unquote(line)
	// 	if err != nil {
	// 		fmt.Println("error converting Unquote")
	// 		fmt.Println(line)
	// 		return
	// 	}
	// 	sumText += len(formatted)
	//
	// }
	
	//encode 
	for scanner.Scan() {
		line := scanner.Text()
		lineLeng := len(line)
		if lineLeng == 0 {
			continue
		}
		sumTotal += lineLeng
		formatted := strconv.Quote(line)
		sumText += len(formatted)
	}

	fmt.Println(sumTotal - sumText)

}

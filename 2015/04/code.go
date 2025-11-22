package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	key := "iwrupvqb"
	lowest := 0

	for true {
		lString := strconv.Itoa(lowest)
		input := key + lString
		hash := md5.Sum([]byte(input))
		hexStr := hex.EncodeToString(hash[:])

		fmt.Println(hexStr)
		fmt.Println(lowest)
		for i := 0; i < 6; i++ {
			if hexStr[i] != '0' {
				lowest++
				break
			}
			if i == 5 {
				fmt.Printf("%d has %d leading zeros\n", i, lowest)
				return
			}
		}
	}

}

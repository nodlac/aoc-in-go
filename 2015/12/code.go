package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	run()
}

func run() {
	file := "./input_user.json"
	// file := "./input_example.json"
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var data any

	err = json.Unmarshal(text, &data)
	if err != nil {
		panic(err)
	}

	sum :=  findNumbers("", data)
	fmt.Println(sum)

}

func findNumbers(path string, value any) float64 {
	var sum float64
	switch v := value.(type) {
	case map[string]any:
		for _, val := range v {
			if val == "red" {
				return 0
			}
		}
		for key, val := range v {
			newPath := key
			if key != "" {
				newPath = path + "." + newPath
			}
			sum += findNumbers(newPath, val)

		}
	case []any:
		for i, val := range v {
			newPath := fmt.Sprintf("%v[%d]", path, i)
			sum += findNumbers(newPath, val)
		}
	case float64:
		sum += v
	}

	return sum
}

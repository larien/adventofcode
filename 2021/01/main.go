package main

import (
	"fmt"

	"github.com/larien/adventofcode/input"
)

const inputFile = "input.txt"

func main() {
	input := input.ParseInt(inputFile)

	result := solve(input)
	fmt.Printf("Result: %v\n", result)
}

func solve(input []int) (result int) {
	var previousDepth *int
	var increases int

	var resultList []int
	for i := range input {
		if i+2 >= len(input) {
			break
		}
		sum := input[i] + input[i+1] + input[i+2]
		resultList = append(resultList, sum)
	}

	for _, d := range resultList {
		dCopy := d
		if previousDepth != nil && d > *previousDepth {
			increases++
		}
		previousDepth = &dCopy
	}

	return increases
}

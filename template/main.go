package main

import (
	"fmt"

	"github.com/larien/adventofcode/input"
)

func main() {
	input := input.ParseString("input.txt")

	result := solve1(input)
	fmt.Printf("Result 1: %v\n", result)

	result = solve2(input)
	fmt.Printf("Result2  .: %v\n", result)
}

func solve1(input []string) (result int) {
	return 0
}

func solve2(input []string) (result int) {
	return 0
}

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/larien/adventofcode/input"
)

func main() {
	input := input.ParseString("input.txt")

	result := solve1(input)
	fmt.Printf("Result 1: %v\n", result)

	result = solve2(input)
	fmt.Printf("Result 2: %v\n", result)
}

//   a
// b e c
//   d
// a = e - 1j
// b = e - 1i
// c = e + 1j
// d = e + 1i
func solve1(input []string) (result int) {
	parsedInput := parseInput(input)

	var lowests []int
	for i := 0; i < len(parsedInput[0]); i++ {
		for j := 0; j < len(parsedInput); j++ {
			var adjacents []int
			value := parsedInput[j][i]
			if i-1 >= 0 {
				adjacents = append(adjacents, parsedInput[j][i-1])
			}
			if i+1 < len(parsedInput[0]) {
				adjacents = append(adjacents, parsedInput[j][i+1])
			}
			if j-1 >= 0 {
				adjacents = append(adjacents, parsedInput[j-1][i])
			}
			if j+1 < len(parsedInput) {
				adjacents = append(adjacents, parsedInput[j+1][i])
			}
			isLower := true
			for _, adjacent := range adjacents {
				if value >= adjacent {
					isLower = false
					break
				}
			}
			if isLower {
				lowests = append(lowests, value)
				fmt.Println(value, adjacents)
			}
		}
	}

	risk := 0
	for _, lowest := range lowests {
		risk += (lowest + 1)
	}
	return risk
}

func solve2(input []string) (result int) {
	return 0
}

func parseInput(input []string) [][]int {
	parsedInput := make([][]int, len(input))

	for i, line := range input {
		inputLine := strings.Split(line, "")
		for _, char := range inputLine {
			n, _ := strconv.Atoi(char)
			parsedInput[i] = append(parsedInput[i], n)
		}
	}
	return parsedInput
}

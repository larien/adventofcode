package main

import (
	"fmt"
	"sort"
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

	lowests := findLowests(parsedInput)

	risk := 0
	for _, lowest := range lowests {
		risk += (lowest[0] + 1)
	}
	return risk
}

func findLowests(input [][]int) [][3]int {
	var positions [][3]int // [ value ] [ i ] [ j ]
	for i := 0; i < len(input[0]); i++ {
		for j := 0; j < len(input); j++ {
			var adjacents []int
			value := input[j][i]
			if i-1 >= 0 {
				adjacents = append(adjacents, input[j][i-1])
			}
			if i+1 < len(input[0]) {
				adjacents = append(adjacents, input[j][i+1])
			}
			if j-1 >= 0 {
				adjacents = append(adjacents, input[j-1][i])
			}
			if j+1 < len(input) {
				adjacents = append(adjacents, input[j+1][i])
			}
			isLower := true
			for _, adjacent := range adjacents {
				if value >= adjacent {
					isLower = false
					break
				}
			}
			if isLower {
				positions = append(positions, [3]int{value, i, j})
			}
		}
	}
	return positions
}

func solve2(input []string) (result int) {
	parsedInput := parseInput(input)

	lowests := findLowests(parsedInput)

	currentLowests := [3]int{9999, 9, 9}
	currentLowest := lowests[0][0]
	for _, lowest := range lowests {
		basin := findBasin(parsedInput, lowest)
		if currentLowest <= basin {
			continue
		}
		currentLowest = addLowests(currentLowests[:], basin)
	}

	mult := 1
	for _, lowest := range currentLowests {
		mult *= lowest
	}

	return mult
}

func findBasin(input [][]int, lowest [3]int) (basin int) {
	return countBasin(input, lowest[1], lowest[2], 0)
}

func countBasin(input [][]int, i, j, basin int) int {
	basin++
	if j-1 >= 0 {
		basin = countBasin(input, i, j-1, basin)
	}
	if j+1 < len(input[0])-1 {
		basin = countBasin(input, i, j+1, basin)
	}
	if i-1 > 0 {
		basin = countBasin(input, i-1, j, basin)
	}
	if i+1 < len(input)-1 {
		basin = countBasin(input, i+1, j, basin)
	}
	return basin
}

func addLowests(currentLowests []int, input int) (lowest int) {
	sort.Ints(currentLowests)

	for i, lowest := range currentLowests {
		if input >= lowest {
			continue
		}
		currentLowests[i] = input
		break
	}

	return currentLowests[0]
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

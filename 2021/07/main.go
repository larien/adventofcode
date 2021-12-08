package main

import (
	"fmt"
	"sort"

	"github.com/larien/adventofcode/input"
	"github.com/larien/adventofcode/utils"
)

func main() {
	input := input.ParseString("input.txt")

	result := solve1(input)
	fmt.Printf("Result 1: %v\n", result)

	result = solve2(input)
	fmt.Printf("Result 2: %v\n", result)
}

func solve1(i []string) (fuel int) {
	positions := input.ParseIntoSlice(i)

	sort.Ints(positions)
	average := utils.Average(positions)
	for position := range positions {
		fuel += abs(average - positions[position])
	}

	return fuel
}

func solve2(i []string) int {
	positions := input.ParseIntoSlice(i)

	sort.Ints(positions)
	candidates := getCandidates(positions)
	var costs []int
	for _, candidate := range candidates {
		referenceCost := calculateCost(positions, candidate)
		costs = append(costs, referenceCost)
	}

	return getLowestCost(costs)
}

// getCandidates returns all possible candidates between the lowest to the highest position
func getCandidates(positions []int) []int {
	var candidates []int
	min, max := positions[0], positions[len(positions)-1]
	for i := min; i <= max; i++ {
		if i < 0 {
			continue
		}
		candidates = append(candidates, i)
	}
	return candidates
}

// calculateCost calculates the cost of the fuel by applying a Gauss summation in the normal cost
// e.g. to go from 1 to 5 the normal cost is 5 - 1 = 4, so the final cost is 1 + 2 + 3 + 4 = 10
func calculateCost(positions []int, reference int) int {
	var result int
	for _, pos := range positions {
		normalCost := abs(reference - pos)
		cost := gaussSum(normalCost)
		result += cost
	}
	return result
}

func getLowestCost(costs []int) int {
	lowestCost := costs[len(costs)-1]
	for _, cost := range costs {
		if lowestCost > cost {
			lowestCost = cost
		}
	}
	return lowestCost
}

// https://letstalkscience.ca/educational-resources/backgrounders/gauss-summation
func gaussSum(an int) int {
	a1 := an - (an - 1)
	return ((a1 + an) * an) / 2
}

// there is a math.Abs, but it's for floats, so it's no biggie to write my own
func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

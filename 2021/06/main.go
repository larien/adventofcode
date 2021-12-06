package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/larien/adventofcode/input"
)

const (
	days1      = 80
	days2      = 256
	newTimer   = 6
	extraTimer = 8
)

func main() {
	input := input.ParseString("input.txt")

	result := solve1(input)
	fmt.Printf("Result 1: %v\n", result)

	result = solve2(input)
	fmt.Printf("Result 2: %v\n", result)
}

func solve1(input []string) (result int) {
	return simulateLanternFish(input, days1)
}

func solve2(input []string) (result int) {
	return simulateLanternFish(input, days2)
}

func simulateLanternFish(input []string, days int) int {
	list := parseInput(input)

	for i := 0; i < days; i++ {
		buffer := make(map[int]int)
		for internalTimer, count := range list {
			if internalTimer == 0 {
				buffer[extraTimer] += count
				buffer[newTimer] += count
				continue
			}
			buffer[internalTimer-1] += count
		}
		list = buffer
	}

	lanternFishCount := 0
	for _, count := range list {
		if count > 0 {
			lanternFishCount += count
		}
	}

	return lanternFishCount
}

func parseInput(input []string) map[int]int {
	result := make(map[int]int)

	resultStr := strings.Split(input[0], ",")
	for _, value := range resultStr {
		i, _ := strconv.Atoi(value)
		result[i]++
	}
	return result
}

// OTHER ATTEMPTS BELOW

// func solve1(input []string) (result int) {
// 	list := parseInput(input)

// 	for i := 0; i < days1; i++ {
// 		var extraTimers []int
// 		for i := range list {
// 			if list[i] == 0 {
// 				list[i] = newTimer
// 				extraTimers = append(extraTimers, extraTimer)
// 				continue
// 			}
// 			list[i]--
// 		}
// 		list = append(list, extraTimers...)
// 	}
// 	lanternFishCount := len(list)

// 	return lanternFishCount
// }

// func parseInput(input []string) []int {
// 	var result []int
// 	resultStr := strings.Split(input[0], ",")
// 	for _, value := range resultStr {
// 		i, err := strconv.Atoi(value)
// 		if err != nil {
// 			panic(err)
// 		}
// 		result = append(result, int(i))
// 	}
// 	return result
// }

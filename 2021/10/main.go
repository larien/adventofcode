package main

import (
	"fmt"
	"strings"

	"github.com/larien/adventofcode/input"
)

var relative = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

func main() {
	input := input.ParseString("input.txt")

	result := solve1(input)
	fmt.Printf("Result 1: %v\n", result)

	result = solve2(input)
	fmt.Printf("Result 2: %v\n", result)
}

func solve1(input []string) (result int) {
	var chars []string
	for _, line := range input {
		_, character := parseInput(strings.Split(line, ""))
		chars = append(chars, character)
	}

	sum := 0
	for _, char := range chars {
		switch char {
		case ")":
			sum += 3
		case "]":
			sum += 57
		case "}":
			sum += 1197
		case ">":
			sum += 25137
		}
	}
	return sum
}

func solve2(input []string) (result int) {
	return 0
}

func parseInput(chars []string) (map[string]int, string) {
	starting := map[string]int{
		"(": 0,
		"[": 0,
		"{": 0,
		"<": 0,
	}
	ending := map[string]int{
		")": 0,
		"]": 0,
		"}": 0,
		">": 0,
	}
	for _, char := range chars {
		if _, ok := starting[char]; ok {
			starting[char]++
			continue
		}
		if _, ok := ending[char]; !ok {
			continue
		}
		ending[char]++
		if starting[relative[char]] != ending[char] {
			return nil, char
		}
	}

	return nil, ""
}

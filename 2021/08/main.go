package main

import (
	"fmt"
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

func solve1(input []string) (result int) {
	entries := parseInput(input)

	uniqueSignalsCount := 0
	for _, entry := range entries {
		for _, output := range entry.outputs {
			switch len(output) {
			case 2, 3, 4, 7:
				uniqueSignalsCount++
			}
		}
	}

	return uniqueSignalsCount
}

func solve2(input []string) (result int) {
	// entries := parseInput(input)

	// segments := parseSegments(renderedSegments)

	// determine all of the wire/segment connections
	// decode the four-digit output values

	return 0
}

func parseSegments(renderedSegments map[int]string) map[int][]string {
	parsedSegments := make(map[int][]string)
	for key, segment := range renderedSegments {
		for _, digit := range validDigits {
			if !strings.Contains(segment, digit) {
				continue
			}
			parsedSegments[key] = append(parsedSegments[key], newMapping[digit])
		}
	}
	return parsedSegments
}

type entry struct {
	inputs  []string
	outputs []string
}

func parseInput(input []string) []entry {
	entries := make([]entry, len(input))
	for i, line := range input {
		parts := strings.Split(line, "|")
		inputs := strings.Split(parts[0], " ")
		outputs := strings.Split(parts[1], " ")
		entries[i] = entry{
			inputs:  inputs[:len(inputs)-1], // remove space
			outputs: outputs[1:],            // remove space
		}
	}
	return entries
}

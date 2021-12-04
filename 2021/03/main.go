package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/larien/adventofcode/input"
)

const inputFile = "input.txt"

func main() {
	input := input.ParseString(inputFile)

	result := solve1(input)
	fmt.Printf("Result: %v\n", result)

	result = solve2(input)
	fmt.Printf("Result: %v\n", result)
}

func solve1(input []string) (result int) {

	invertedBits := make([][]string, len(input[0]))
	for _, bits := range input {
		for i, bit := range bits {
			invertedBits[i] = append(invertedBits[i], string(bit))
		}
	}

	gammaRate := make([]string, len(input[0]))
	for _, bits := range invertedBits {
		count0, count1 := 0, 0
		for _, bit := range bits {
			if bit == "0" {
				count0++
				continue
			}
			count1++
		}
		if count0 > count1 {
			gammaRate = append(gammaRate, "0")
			continue
		}
		gammaRate = append(gammaRate, "1")
	}

	epsilonRate := make([]string, len(input[0]))
	for _, bit := range gammaRate {
		if bit == "0" {
			epsilonRate = append(epsilonRate, "1")
			continue
		}
		epsilonRate = append(epsilonRate, "0")
	}

	gamma, _ := strconv.ParseInt(strings.Join(gammaRate, ""), 2, 64)
	epsilon, _ := strconv.ParseInt(strings.Join(epsilonRate, ""), 2, 64)

	return int(gamma) * int(epsilon)
}

func solve2(input []string) (result int) {
	inputOxygen := make([]string, len(input))
	copy(inputOxygen, input)

	index := 0
	for len(inputOxygen) > 1 {
		inputOxygen = processO(inputOxygen, index)
		index++
	}

	index = 0
	for len(input) > 1 {
		input = processC(input, index)
		index++
	}

	oxygen, _ := strconv.ParseInt(inputOxygen[0], 2, 64)
	co2, _ := strconv.ParseInt(input[0], 2, 64)

	return int(oxygen) * int(co2)
}

func processO(oxygen []string, index int) []string {
	count0, count1 := getBitsCount(oxygen, index)

	mostCommon := "0"
	if count1 > count0 || count0 == count1 {
		mostCommon = "1"
	}

	var output []string
	for _, bits := range oxygen {
		if string(bits[index]) == mostCommon {
			output = append(output, bits)
		}
	}

	return output
}

func processC(co2 []string, index int) []string {
	count0, count1 := getBitsCount(co2, index)

	leastCommon := "1"
	if count1 > count0 || count0 == count1 {
		leastCommon = "0"
	}

	var output []string
	for _, bits := range co2 {
		if string(bits[index]) == leastCommon {
			output = append(output, bits)
		}
	}

	return output
}

func getBitsCount(input []string, index int) (count0, count1 int) {
	for _, bits := range input {
		if string(bits[index]) == "0" {
			count0++
			continue
		}
		count1++
	}
	return count0, count1
}

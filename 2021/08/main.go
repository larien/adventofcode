package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/larien/adventofcode/input"
	"github.com/larien/adventofcode/utils"
)

var uniqueSignalsLength = map[int]int{
	// length: number
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

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
	entries := parseInput(input)

	sum := 0
	for _, entry := range entries {
		decoded := decode(entry.inputs)
		sum += calculateOutput(decoded, entry.outputs)
	}
	return sum
}

func calculateOutput(decoded map[int]string, outputs []string) int {
	var result []string
	for _, output := range outputs {
		for number, letters := range decoded {
			x := strings.Split(output, "")
			sort.Strings(x)

			y := strings.Split(letters, "")

			if strings.Join(x, "") != strings.Join(y, "") {
				continue
			}
			result = append(result, fmt.Sprint(number))
		}
	}

	value, _ := strconv.Atoi(strings.Join(result, ""))
	return value
}

//  1111
// 2    3
// 2    3
//  4444
// 5    6
// 5    6
//  7777
// 1: 3, 6
// 2: 1, 3, 4, 5, 7
// 3: 1, 3, 4, 6, 7
// 4: 2, 3, 4, 6
// 5: 1, 2, 4, 6, 7
// 6: 1, 2, 4, 5, 6, 7
// 7: 1, 3, 6
// 8: 1, 2, 3, 4, 5, 6, 7
// 9: 1, 2, 3, 4, 6, 7

type segments map[int][]string

func decode(pattern []string) map[int]string {
	decoded := make(segments, len(pattern))
	mapping := make(map[int]string, 7)

	pattern = decoded.decodeUnique(pattern)
	mapping[1] = getDifference(decoded[7], decoded[1])[0]
	pattern = decoded.number9(mapping, pattern)
	mapping[5] = getDifference(decoded[8], decoded[9])[0]
	mapping[7] = decoded.segment7(mapping, pattern)
	pattern = decoded.number2(mapping, pattern)
	pattern = decoded.number3and5(mapping, pattern)
	pattern = decoded.number0and6(mapping, pattern)

	if len(pattern) != 0 {
		log.Fatal("pattern not empty", pattern)
	}

	output := make(map[int]string, len(decoded))
	for k, v := range decoded {
		sort.Strings(v)
		output[k] = strings.Join(v, "")
	}
	return output
}

func (s segments) segment7(mapping map[int]string, pattern []string) string {
	for _, segment := range pattern {
		if len(segment) == 6 {
			// 6 or 0
			// remove 4, segments 1 and 5 and segment 7 is left
			differences := getDifference(strings.Split(segment, ""), s[4])
			differences = utils.RemoveFromSlice(differences, mapping[1])
			differences = utils.RemoveFromSlice(differences, mapping[5])
			return differences[0]
		}
	}
	return ""
}

func (s segments) decodeUnique(pattern []string) []string {
	var newPattern []string
	for _, segment := range pattern {
		switch len(segment) {
		case 2, 3, 4, 7:
			s[uniqueSignalsLength[len(segment)]] = strings.Split(segment, "")
			continue
		default:
			newPattern = append(newPattern, segment)
		}
	}
	return newPattern
}

func (s segments) number3and5(mapping map[int]string, pattern []string) []string {
	seg3 := ""
	seg5 := ""
	for _, segment := range pattern {
		if len(segment) != 5 {
			// 3 and 5 have 5 segments
			continue
		}
		differences := getDifference(s[2], strings.Split(segment, ""))
		if len(differences) == 2 {
			// 2 - 5 = 0 segments
			seg5 = segment
			continue
		}
		if len(differences) == 1 {
			// 2 - 3 = 1 segment
			seg3 = segment
			continue
		}
	}

	s[3] = strings.Split(seg3, "")
	s[5] = strings.Split(seg5, "")
	pattern = utils.RemoveFromSlice(pattern, seg3)
	return utils.RemoveFromSlice(pattern, seg5)
}

func (s segments) number0and6(mapping map[int]string, pattern []string) []string {
	seg0 := ""
	seg6 := ""
	for _, segment := range pattern {
		if len(segment) != 6 {
			// 0 and 6 have 6 segments
			continue
		}
		differences := getDifference(strings.Split(segment, ""), s[5])
		if len(differences) == 1 {
			// 5 - 6 = 1 segment
			seg6 = segment
			continue
		}
		if len(differences) == 2 {
			// 5 - 0 = 2 segments
			seg0 = segment
			continue
		}
	}

	s[0] = strings.Split(seg0, "")
	s[6] = strings.Split(seg6, "")
	pattern = utils.RemoveFromSlice(pattern, seg0)
	return utils.RemoveFromSlice(pattern, seg6)
}

func (s segments) number9(mapping map[int]string, pattern []string) []string {
	// segment 1 + number 4 = number 9
	letters := s[4]
	letters = append(letters, mapping[1])

	seg := ""
	for _, segment := range pattern {
		if len(segment) != 6 {
			// 9 has 6 segments
			continue
		}
		contains := true
		for _, letter := range letters {
			if !strings.Contains(segment, letter) {
				contains = false
			}
		}
		if contains {
			seg = segment
			break
		}
	}

	s[9] = strings.Split(seg, "")
	return utils.RemoveFromSlice(pattern, seg)
}

func (s segments) number2(mapping map[int]string, pattern []string) []string {
	seg := ""
	for _, segment := range pattern {
		if len(segment) == 5 {
			// number 2 less number 4 has 3 segments, 3 and 5 has 2
			differences := getDifference(strings.Split(segment, ""), s[4])
			if len(differences) == 2 {
				continue
			}
			seg = segment
			break
		}
	}

	s[2] = strings.Split(seg, "")
	return utils.RemoveFromSlice(pattern, seg)
}

func (s segments) number6(mapping map[int]string, pattern []string) []string {
	seg := ""
	for _, segment := range pattern {
		if len(segment) == 6 {
			// number 6 has 6 segments besides number 9
			seg = segment
			break
		}
	}

	s[6] = strings.Split(seg, "")
	return utils.RemoveFromSlice(pattern, seg)
}

func (s segments) number5(mapping map[int]string, pattern []string) []string {
	seg := ""

	// number 5 = number 6 - segment 5
	letters := utils.RemoveFromSlice(s[6], mapping[5])

	for _, segment := range pattern {
		contains := true
		for _, letter := range letters {
			if !strings.Contains(segment, letter) {
				contains = false
			}
		}
		if contains {
			seg = segment
			break
		}
	}

	s[5] = strings.Split(seg, "")
	return utils.RemoveFromSlice(pattern, seg)
}

func getDifference(a, b []string) []string {
	var difference []string
	for _, digit := range a {
		if !utils.StringInSlice(b, digit) {
			difference = append(difference, digit)
		}
	}
	return difference
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

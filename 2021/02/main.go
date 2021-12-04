package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/larien/adventofcode/input"
)

const (
	inputFile = "input.txt"

	forward = "forward"
	down    = "down"
	up      = "up"
	aim     = "aim"
)

func main() {
	input := input.ParseString(inputFile)

	result := solve1(input)
	fmt.Printf("Result 1: %v\n", result)

	result = solve2(input)
	fmt.Printf("Result 2: %v\n", result)
}

func solve1(input []string) (result int) {
	horizontal, depth := 0, 0

	for _, line := range input {
		lineSplit := strings.Split(line, " ")
		direction := lineSplit[0]
		value, _ := strconv.Atoi(lineSplit[1])

		switch direction {
		case forward:
			horizontal += value
		case down:
			depth += value
		case up:
			depth -= value
		default:
			log.Fatal("Unknown direction")
		}
	}

	return horizontal * depth
}

func solve2(input []string) (result int) {
	commands := parseInput(input)
	var horizontal, depth, aim int

	for _, c := range commands {
		command := c
		switch command.direction {
		case forward:
			horizontal += command.value
			if aim == 0 {
				continue
			}
			depth += aim * command.value
		case down:
			aim += command.value
		case up:
			aim -= command.value
		default:
			log.Println("Unknown direction:", command.direction)
		}
	}

	return horizontal * depth
}

type command struct {
	direction string
	value     int
}

func parseInput(input []string) []command {
	commands := make([]command, len(input))
	for _, line := range input {
		var command command
		lineSplit := strings.Split(line, " ")
		command.direction = lineSplit[0]
		command.value, _ = strconv.Atoi(lineSplit[1])

		commands = append(commands, command)
	}
	return commands
}

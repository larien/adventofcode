package main

import (
	"fmt"

	"github.com/larien/adventofcode/input"
)

const (
	diagramSize = 1000
)

type vent struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	input := input.ParseString("input.txt")

	result := solve1(input)
	fmt.Printf("Result 1: %v\n", result)

	result = solve2(input)
	fmt.Printf("Result 2: %v\n", result)
}

type diagram [diagramSize][diagramSize]int

func solve1(input []string) (result int) {
	vents := parseInput(input)

	var diagram diagram
	for _, vent := range vents {
		isDiagonal := vent.x1 != vent.x2 && vent.y1 != vent.y2
		if isDiagonal {
			continue
		}
		diagram.mark(vent)
	}

	return diagram.countOverlap()
}

func solve2(input []string) (result int) {
	vents := parseInput(input)

	var diagram diagram
	for _, vent := range vents {
		isDiagonal := vent.x1 != vent.x2 && vent.y1 != vent.y2
		if isDiagonal {
			diagram.markDiagonal(vent)
			continue
		}
		diagram.mark(vent)
	}

	return diagram.countOverlap()
}

func (d diagram) countOverlap() int {
	count := 0
	for _, row := range d {
		for _, value := range row {
			if value > 1 {
				count++
			}
		}
	}
	return count
}

func (d *diagram) mark(vent vent) {
	x, y := vent.x1, vent.y1
	for {
		d[y][x]++
		if x == vent.x2 && y == vent.y2 {
			break
		}
		if x != vent.x2 {
			x++
		}
		if y != vent.y2 {
			y++
		}
	}
}

func parseInput(input []string) (vents []vent) {
	for _, line := range input {
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if x1 > x2 || y1 > y2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}

		vents = append(vents, vent{x1, y1, x2, y2})
	}
	return
}

func (d *diagram) markDiagonal(vent vent) {
	x, y := vent.x1, vent.y1
	for {
		d[y][x]++
		if x == vent.x2 && y == vent.y2 {
			break
		}
		if x != vent.x2 {
			if vent.x1 < vent.x2 {
				x++
			} else {
				x--
			}
		}
		if y != vent.y2 {
			if vent.y1 < vent.y2 {
				y++
			} else {
				y--
			}
		}
	}
}

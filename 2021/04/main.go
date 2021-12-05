package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/larien/adventofcode/input"
)

const inputFile = "input.txt"

type board [5][5]int
type boards []board

func main() {
	input := input.ParseString(inputFile)

	result := solve1(input)
	fmt.Printf("Result 1: %v\n", result)

	result = solve2(input)
	fmt.Printf("Result 2: %v\n", result)
}

func solve1(input []string) (result int) {
	numbers, boards := parseInput(input)

	winnerIndex := 0
	winnerNumber := 0
	for _, number := range numbers {
		for i := range boards {
			boards[i].markNumber(number)
		}
		winners := boards.checkWinners()
		if len(winners) > 0 {
			winnerNumber = number
			winnerIndex = winners[0]
			break
		}
	}

	score := boards[winnerIndex].calculateScore()

	return score * winnerNumber
}

func solve2(input []string) (result int) {
	numbers, boards := parseInput(input)

	winnerNumber := 0
	lastIndex := 0
	winnerSnapshots := make(map[int]board)
	for _, number := range numbers {
		for i := range boards {
			boards[i].markNumber(number)
		}
		winners := boards.checkWinners()
		if len(winners) > 0 {
			for _, winnerIndex := range winners {
				if _, ok := winnerSnapshots[winnerIndex]; !ok {
					winnerSnapshots[winnerIndex] = boards[winnerIndex]
					lastIndex = winnerIndex
				}
			}
			if len(winnerSnapshots) == len(boards) {
				winnerNumber = number
				break
			}
		}
	}

	score := winnerSnapshots[lastIndex].calculateScore()
	return score * winnerNumber
}

func (b *board) markNumber(number int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			boardNumber := b[i][j]
			if boardNumber == number {
				b[i][j] = -1
			}
		}
	}
}

func (b boards) checkWinners() map[int]int {
	winners := make(map[int]int)
	count := 0
	for i, board := range b {
		if board.checkRow() {
			winners[count] = i
			count++
			continue
		}
		if board.checkColumn() {
			winners[count] = i
			count++
		}
	}

	return winners
}

func (b board) checkRow() bool {
	for i := 0; i < 5; i++ {
		winner := true
		for j := 0; j < 5; j++ {
			number := b[i][j]
			if number != -1 {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}
	return false
}

func (b board) checkColumn() bool {
	for j := 0; j < 5; j++ {
		winner := true
		for i := 0; i < 5; i++ {
			number := b[i][j]
			if number != -1 {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}
	return false
}

func (b board) calculateScore() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] != -1 {
				sum += b[i][j]
			}
		}
	}
	return sum
}

func parseInput(input []string) (numbers []int, boards boards) {
	numbersStr := strings.Split(input[0], ",")
	for _, numberStr := range numbersStr {
		number, _ := strconv.Atoi(numberStr)
		numbers = append(numbers, number)
	}

	boardList := input[2:]
	boardRow := 0

	var b board
	for _, row := range boardList {
		if len(row) == 0 {
			boards = append(boards, b)
			boardRow = 0
			b = board{}
			continue
		}

		boardNumbers := strings.Split(row, " ")
		i := 0
		for _, number := range boardNumbers {
			if number == "" {
				continue
			}
			b[boardRow][i], _ = strconv.Atoi(number)
			i++
		}
		boardRow++
	}
	boards = append(boards, b)

	return numbers, boards
}

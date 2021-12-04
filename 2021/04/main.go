package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/larien/adventofcode/input"
)

// TODO: 7 and 75 wins, but only 7 counts since it's the first one

const inputFile = "input.txt"

type board [5][5]int
type boards []board

func main() {
	input := input.ParseString(inputFile)

	result := solve1(input)
	fmt.Printf("Result: %v\n", result)

	result = solve2(input)
	fmt.Printf("Result: %v\n", result)
}

func solve1(input []string) (result int) {

	numbers, boards := parseInput(input)

	winnerIndex := 0
	winnerNumber := 0
	for _, number := range numbers {
		fmt.Println("number: ", number)
		for i := range boards {
			boards[i].markNumber(number)
			// boards[i].printBoard()
		}
		var winner bool
		winnerIndex, winner = boards.checkWinners()
		if winner {
			winnerNumber = number
			break
		}
		// time.Sleep(time.Second * 1)
	}

	score := boards[winnerIndex].calculateScore()
	fmt.Println("winner number: ", winnerNumber)
	fmt.Println("winner board: ", winnerIndex)
	return score * winnerNumber
}

func (b board) printBoard() {
	fmt.Println("-------------------")
	for i := 0; i < 5; i++ {
		fmt.Println(b[i])
	}
	fmt.Println("-------------------")
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

func (b boards) checkWinners() (int, bool) {
	for i, board := range b {
		if board.checkRow() {
			return i, true
		}
		if board.checkColumn() {
			return i, true
		}
	}

	return 0, false
}

func (b board) checkRow() bool {
	winner := true
	for i := 0; i < 5; i++ {
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
	winner := true
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if b[j][i] != -1 {
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

func solve2(input []string) (result int) {
	return 0
}

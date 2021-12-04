package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ParseInt return content of file as int
func ParseInt(path string) []int {
	result := make([]int, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

// ParseString return content of file as string
func ParseString(path string) []string {
	result := make([]string, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

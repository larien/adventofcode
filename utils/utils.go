package utils

import "sort"

func Average(numbers []int) int {
	return numbers[len(numbers)/2]
}

func Median(numbers []int) int {
	sort.Ints(numbers)

	index := len(numbers) / 2
	if index%2 == 0 {
		return numbers[index]
	}

	return (numbers[index-1] + numbers[index]) / 2
}

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

func StringInSlice(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func RemoveFromSlice(s []string, value string) []string {
	out := make([]string, len(s))
	i := 0
	for _, v := range s {
		if v != value {
			out[i] = v
			i++
		}
	}
	return out[:i]
}

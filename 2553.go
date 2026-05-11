package main

import "strconv"

func separateDigits(nums []int) []int {
	result := []int{}
	arr := []rune{}

	for _, num := range nums {
		arr = []rune(strconv.Itoa(num))
		for _, char := range arr {
			result = append(result, int(char - '0'))
		}
	}

	return result
}
package main

import (
	"math"
)

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	result := 0
	grid := make(map[int]int)

	for _, num1 := range arr1 {
		for num1 >= 1 {
			grid[num1] = 1
			num1 /= 10
		}
	}

	for _, num2 := range arr2 {
		for num2 >= 1 {
			_, ok := grid[num2]
			if ok {
				result = int(math.Max(float64(result), (math.Log10(float64(num2)) + 1)))
				break
			}
			num2 /= 10
		}
	}
	

	return result
}
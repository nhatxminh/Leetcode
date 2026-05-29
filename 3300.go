package main

import "math"

func minElement(nums []int) int {
	result := math.MaxInt

	for _, num := range nums {
		current := 0
		for num >= 1 {
			current += num % 10
			num /= 10
		}

		result = int(math.Min(float64(result), float64(current)))
	}

	return result
}

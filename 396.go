package main

import "math"

func maxRotateFunction(nums []int) int {
	prev := 0
  sum := 0
	result := math.MinInt
	
	for idx, value := range nums {
		sum += value
		prev += idx * value
	}

	for i := range nums {
		if i == 0 {
			result = int(math.Max(float64(prev), float64(result)))
			continue
		}
		
		prev = prev + sum - len(nums) * nums[len(nums) - i]
		result = int(math.Max(float64(prev), float64(result)))
	}

	return result
}
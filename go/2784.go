package main

import (
	"math"
)

func isGood(nums []int) bool {
	max := int(math.MinInt)
	grid := make(map[int]int)

	for _, num := range nums {
		max = int(math.Max(float64(max), float64(num)))
		value, ok := grid[num]
		if ok {
			grid[num] = value + 1
		} else {
			grid[num] = 1
		}
	}

	if len(nums) != max + 1 {
		return false
	}

	for i := 1; i <= max; i++  {
		value, ok := grid[i]
		if ok {
			if i == max {
				if value != 2 {
					return false
				}
			} else {
				if value == 2 {
					return false
				} 
			}
		} else {
			return false
		}
	}

	return true
}
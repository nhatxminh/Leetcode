package main

import "math"

func maxDistance(colors []int) int {
	result := 0

	for i := len(colors) - 1; i > 0; i-- {
		if colors[i] != colors[0] {
			result = int(math.Max(float64(result), float64(i)))
			break
		}
	}

	for j := 0; j < len(colors) - 1; j++ {
		if colors[j] != colors[len(colors) - 1] {
			result = int(math.Max(float64(result), float64(len(colors) - 1 - j)))
			break
		}
	}

	return result
}
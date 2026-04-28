package main

import (
	"fmt"
	"math"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	result := 0
	arr := []int{}

	for _, outer := range grid {
		for _, inner := range outer {
			arr = append(arr, inner)
		}
	}

	sort.Ints(arr)
	fmt.Println(arr)
	idx := len(arr) / 2
	compare := arr[idx]
	fmt.Println(compare)

	for i := idx; i < len(arr); i++ {
		if compare - arr[i] == 0 {
			continue
		} else if (compare - arr[i]) % x == 0 {
			result += int(math.Abs(float64((compare - arr[i]) / x)))
		} else {
			return -1
		}
	}

	for i := idx; i >= 0; i-- {
		if compare - arr[i] == 0 {
			continue
		} else if (compare - arr[i]) % x == 0 {
			result += int(math.Abs(float64((compare - arr[i]) / x)))
		} else {
			return -1
		}
	}

	return result
}
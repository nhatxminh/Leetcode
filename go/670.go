package main

import (
	"math"
	"sort"
)

func maximumSwap(num int) int {
	length := int(math.Floor(math.Log10(float64(num)))) + 1
	position := length - 1
	list := make([]int, length)
	curr := 0
	grid := make(map[int]int)
	result := 0

	for num >= 1 {
		curr = num % 10
		list[position] = curr
		_, ok := grid[curr]
		if !ok {
			grid[curr] = position
		}
		position--
		num /= 10
	}

	sortedList := make([]int, len(list))
	copy(sortedList, list)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedList)))

	for idx, num := range sortedList {
		if list[idx] != num {
			temp := list[idx]
			list[idx] = num
			swap, ok := grid[num]
			if ok {
				list[swap] = temp
			}
			break
		}
	}

	for idx, num := range list {
		result += int(math.Pow10(len(list)-idx-1)) * num
	}

	return result
}

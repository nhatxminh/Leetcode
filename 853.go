package main

import (
	"slices"
)

func carFleet(target int, position []int, speed []int) int {
	m := make(map[int]int)
	total := 1
	prev := 0.0
	current := 0.0

	for i := range position {
		m[position[i]] = speed[i]
	}

	slices.Sort(position)

	for i := len(position) - 1; i >= 0; i-- {
		current = (float64(target) - float64(position[i]))/float64(m[position[i]])
		if i == len(position) - 1 {
			prev = current
		}

		if current > prev {
			total++
			prev = current
		}
	}

	return total
}
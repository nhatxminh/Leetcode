package main

import "fmt"

func numberOfWays(s string) int64 {
	result := 0
	first := make([]int, 0)
	second := make([]int, 0)
	firstSum := 0
	secondSum := 0
	count := 0
	// curr := 1
	var prev rune

	for idx, char := range s {
		if idx == 0 {
			prev = char
			count++
			continue
		}

		if char == prev {
			count++
		} else {
			if len(first) == len(second) {
				first = append(first, count)
				firstSum += count
			} else {
				second = append(second, count)
				secondSum += count
			}
			count = 1
			prev = char
		}

		if idx == len(s)-1 {
			if len(first) == len(second) {
				first = append(first, count)
				firstSum += count
			} else {
				second = append(second, count)
				secondSum += count
			}
		}
	}

	fmt.Print("First: ")
	fmt.Println(first)
	fmt.Print("Second: ")
	fmt.Println(second)

	left := 0
	right := firstSum
	for idx, num := range second {
		left += first[idx]
		right -= first[idx]
		result += num * left * right
		fmt.Println(result)
	}

	left = 0
	right = secondSum
	for idx, num := range first {
		if idx < 1 {
			continue
		}
		left += second[idx-1]
		right -= second[idx-1]
		result += num * left * right
		fmt.Println(result)
	}

	return int64(result)
}

package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	arr := []rune(str)

	for i := range len(arr) {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}

	return true
}
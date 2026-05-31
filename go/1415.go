package main

import (
	"fmt"
)

func getHappyString(n int, k int) string {
	result := make([]string, 0)
	iterate := make([]byte, 0)
	s := []byte{'a', 'b', 'c'}

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == n {
			perm := make([]byte, n)
			copy(perm, iterate)
			fmt.Println("Final string: " + string(perm))
			result = append(result, string(perm))
			return
		}

		for i := range s {
			if len(iterate) > 0 && iterate[len(iterate) - 1] == s[i] {
				continue
			}
			iterate = append(iterate, s[i])
			backtrack(start + 1)
			iterate = iterate[:len(iterate) - 1]
		}

	}

	backtrack(0)

	if k <= len(result) {
		return result[k - 1]
	}

	return ""
}
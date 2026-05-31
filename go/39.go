package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	current := []int{}

	var backtrack func(list []int, current []int, remain int)
	backtrack = func(list []int, current []int, remain int) {
		fmt.Print("Current :")
		fmt.Print(current)
		fmt.Print(" Remain: ")
		fmt.Println(remain)
		for _, num := range list {
			if remain == num {
				new := make([]int, len(current))
				copy(new, current)
				fmt.Print("New: ")
				fmt.Println(new)
				new = append(new, num)
				result = append(result, new)
			} else if remain > num {
				current = append(current, num)
				backtrack(list, current, remain-num)
				current = current[:len(current)-1]
			}
			list = list[1:]
		}
	}

	backtrack(candidates, current, target)

	return result
}

package main

func subsets(nums []int) [][]int {
  result := make([][]int, 0)
	var nilSlice []int

	result = append(result, nilSlice)

	for _, num := range nums {
		for _, slice := range result {
			subset := make([]int, len(slice))
			copy(subset, slice)
			subset = append(subset, num)
			result = append(result, subset)
		}
	}

	return result
}
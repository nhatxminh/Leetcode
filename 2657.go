package main 

func findThePrefixCommonArray(A []int, B []int) []int {
  grid := make(map[int]int)
	result := make([]int, len(A))

	for idx, value := range A {
		grid[value] = idx
	}

	for idx, value := range B {
		position, ok := grid[value]
		if ok {
			if idx >= position {
				if idx > 0 {
					result[idx] += result[idx-1]
				}
				result[idx]++
			} else {
				result[position]++
			}
		}
	}

	return result
}
package main

func maxDepthAfterSplit(seq string) []int {
  result := make([]int, len(seq))
	arr := []rune(seq)
	var prev rune

	for idx, char := range arr {
		if idx == 0 {
			prev = char
			continue
		} 

		if prev == char {
			if result[idx - 1] == 0 {
				result[idx] = 1
			} else {
				result[idx] = 0
			}
		}

		if prev !=char {
			prevRune := result[idx - 1]
			result[idx] = prevRune
		}

		prev = char
	}


	return result
}
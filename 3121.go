package main

func numberOfSpecialChars(word string) int {
	result := 0
	grid := make(map[rune]int)

	for _, char := range word {
		_, ok := grid[char]
		if !ok {
			grid[char] = 0
		}

		if char <= 90 {
			freq, ok := grid[char+32]
			if ok && freq < 1 {
				result++
				grid[char+32] = 1
			}
		} else if char >= 97 {
			freq, ok := grid[char-32]
			if ok && freq < 1 {
				grid[char-32] = 1
				if grid[char] == 1 {
					result--
				} else {
					grid[char] = 1
				}
			}

		}
	}

	return result
}

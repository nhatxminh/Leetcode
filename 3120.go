package main

// func numberOfSpecialChars(word string) int {
// 	result := 0
// 	grid := make(map[rune]int)
//
// 	for _, char := range word {
// 		_, ok := grid[char]
// 		if ok {
// 			continue
// 		} else {
// 			grid[char] = 0
// 		}
//
// 		if char > 97 {
// 			freq, ok := grid[char-32]
// 			if ok && freq < 1 {
// 				result++
// 				grid[char-32] = 1
// 			}
// 		} else {
// 			freq, ok := grid[char+32]
// 			if ok && freq < 1 {
// 				result++
// 				grid[char+32] = 1
// 			}
// 		}
// 	}
//
// 	return result
// }

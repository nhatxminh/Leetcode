package main

func romanToInt(s string) int {
  result := 0
	current := 0
	prev := 0

	for _, roman := range s {
		current = convertRomanToInt(roman)
		if prev < current {
			result -= 2 * prev
		}
		result += current
		prev = current 
	}

	return result
}

func convertRomanToInt(roman rune) int {
	switch roman {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
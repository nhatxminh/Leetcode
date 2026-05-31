package main

import "math"

func furthestDistanceFromOrigin(moves string) int {
	l := 0
	r := 0
	blank := 0

	for _, value := range moves {
		switch value {
		case 'L':
			l++
		case 'R':
			r++
		default:
			blank++
		}
	}

	return int(math.Abs(float64(l) - float64(r))) + blank
}
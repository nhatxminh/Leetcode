package main

import "math"

func addBinary(a string, b string) string {
	stringA := []rune(a)
	stringB := []rune(b)
	length := int(math.Max(float64(len(a)), float64(len(b)))) + 1
	result := make([]rune, length)
	position := len(result) - 1
	positionA := len(a) - 1
	positionB := len(b) - 1
	prev := 0
	numA := 0
	numB := 0

	for position >= 0 {
		if positionA >= 0 {
			numA = int(stringA[positionA] - '0')
			positionA--
		} else {
			numA = 0
		}

		if positionB >= 0 {
			numB = int(stringB[positionB] - '0')
			positionB--
		} else {
			numB = 0
		}

		curr := numA + numB

		switch curr + prev {
		case 0:
			prev = 0
			result[position] = '0'
		case 1:
			prev = 0
			result[position] = '1'
		case 2:
			prev = 1
			result[position] = '0'
		case 3:
			prev = 1
			result[position] = '1'
		}
		position--
	}

	if result[0] == '0' {
		result = result[1:]
	}

	return string(result)
}

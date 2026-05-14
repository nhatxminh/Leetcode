package main

import (
	"math"
	"strconv"
)

func mirrorDistance(n int) int {
	new := []rune{}
	arr := []rune(strconv.Itoa(n))

	for i := len(arr) - 1; i >= 0; i-- {
		new = append(new, arr[i])
	}

	str := string(new)
	mirror, _ := strconv.Atoi(str)

	return int(math.Abs(float64(mirror) - float64(n)))
}
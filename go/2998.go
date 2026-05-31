package main

import "fmt"

func minimumOperationsToMakeEqual(x int, y int) int {
	curr := 0
	step := 0
	currWidth := 0
	bfs := make([]int, 0)
	width := make(map[int]int)
	visited := make(map[int]int)

	bfs = append(bfs, x)
	width[0] = 1

	for {
		if bfs[0] == y {
			fmt.Print(width)
			return step
		}

		fmt.Println(bfs)
		curr = bfs[0]
		bfs = bfs[1:]
		visited[curr] = 0
		currWidth++

		_, plus := visited[curr+1]
		if !plus {
			bfs = append(bfs, curr+1)
			_, ok := width[step+1]
			if ok {
				width[step+1] = width[step+1] + 1
			} else {
				width[step+1] = 1
			}
		}

		_, minus := visited[curr-1]
		if !minus {
			bfs = append(bfs, curr-1)
			_, ok := width[step+1]
			if ok {
				width[step+1] = width[step+1] + 1
			} else {
				width[step+1] = 1
			}
		}

		if curr%11 == 0 {
			_, div11 := visited[curr/11]
			if !div11 {
				bfs = append(bfs, curr/11)
				_, ok := width[step+1]
				if ok {
					width[step+1] = width[step+1] + 1
				} else {
					width[step+1] = 1
				}
			}
		}

		if curr%5 == 0 {
			_, div11 := visited[curr/5]
			if !div11 {
				bfs = append(bfs, curr/5)
				_, ok := width[step+1]
				if ok {
					width[step+1] = width[step+1] + 1
				} else {
					width[step+1] = 1
				}
			}
		}

		if currWidth == width[step] {
			step++
			currWidth = 0
		}
	}
}

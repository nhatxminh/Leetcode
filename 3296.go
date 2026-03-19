package main

import (
	"container/heap"
	"slices"
	"math"
)

type intHeap []int

func (h intHeap) Len() int { return len(h) }
func (h intHeap) Less(i, j int) bool {return h[i] < h[j]}
func (h intHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]} 

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	old := *h
	n := len(*h)
	x := old[n - 1]
	*h = old[0: n - 1]

	return x
}

func (h intHeap) Peek() int {
	return h[0]
}

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	var final int
	track := make(map[int][][2]int)

	slices.Sort(workerTimes)
	list := make(intHeap, len(workerTimes))
	copy(list, workerTimes)
	heap.Init(&list)

	for i := range workerTimes {
		_, ok := track[workerTimes[i]] 
		if ok {
			track[workerTimes[i]] = append(track[workerTimes[i]], [2]int{workerTimes[i], 1})
		} else {
			track[workerTimes[i]] = [][2]int{{workerTimes[i], 1}}
		}
	}

	for {
		if mountainHeight <= 0 {
			break
		}

		min := list.Peek()
		value, ok := track[min]
		if ok {
			final = int(math.Max(float64(final), float64(heap.Pop(&list).(int))))
			newValue := final + value[0][0] * (value[0][1] + 1)
			heap.Push(&list, newValue)

			if len(value) > 1 {
				track[min] = value[1:]
			}

			_, ok := track[newValue] 
			if ok {
				track[newValue] = append(track[newValue], [2]int{value[0][0], value[0][1] + 1})
			} else {
				track[newValue] = [][2]int{{value[0][0], value[0][1] + 1}}
			}
			mountainHeight--
		}
	}

	return int64(final)
}



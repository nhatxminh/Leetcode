package main

import (
	"container/heap"
)

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		return pq[i].value > pq[j].value
	}
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)

	item := old[n-1]
	old[n-1] = nil 

	*pq = old[0 : n-1]
	return item
}

func jump(nums []int) int {
	result := 0
	steps := 0
	position := 0
	done := false
	
	var dfs func(position int, steps int)
	
	dfs = func(position int, steps int) {
		h := &PriorityQueue{}
		heap.Init(h)
		if position >= len(nums) - 1 {
			done = true
			result = steps
			return
		}

		if done == true {
			return
		}
	
		for i := range nums[position] {
			if position + i + 1 == len(nums) - 1 {
				result = steps + 1
				done = true
				return
			}
			if position + i + 1 < len(nums) {
				heap.Push(h, &Item{value: position + i + 1, priority: position + i + 1 + nums[position + i + 1]})
			}
		}

		for len(*h) != 0 {
			dfs(heap.Pop(h).(*Item).value, steps + 1)
			if done {
				break
			}
		}
	}

	dfs(position, steps)

	return result
}
package main 

func canReach(arr []int, start int) bool {
	length := len(arr)
	visited := make(map[int]bool, length)

	path := make([]int, 0)
	path = append(path, start)

	for len(path) > 0 {
		curr := path[len(path) - 1]
		path = path[:len(path) - 1]

		if arr[curr] == 0 {
			return true
		}

		forward := curr + arr[curr]
		if forward < length && !visited[forward] {
			visited[forward] = true 
			path = append(path, forward)
		}

		backward := curr - arr[curr] 
		if backward >= 0 && !visited[backward] {
			visited[backward] = true
			path = append(path, backward)
		}
	}

	return false  
}
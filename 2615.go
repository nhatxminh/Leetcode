package main

func distance(nums []int) []int64 {
	prefix := 0
	m := make(map[int][]int)
	result := make([]int64, len(nums))

	for idx, value := range nums {
		arr, ok := m[value]
		if ok {
			arr = append(arr, idx)
			m[value] = arr
		} else {
			arr := []int{}
			arr = append(arr, idx)
			m[value] = arr
		}
	}

	for _, arr := range m {
		if len(arr) < 2 {
			continue
		}

		for i, value := range arr {
			prefix += value
			result[value] += int64((i + 1) * value - prefix)
		}

		prefix = 0

		for i := len(arr) - 1; i >=0; i-- {
			prefix += arr[i]
			result[arr[i]] += int64((len(arr) - i) * arr[i] - prefix) * (-1)
		}

		prefix = 0
	}

	return result
}
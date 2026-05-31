package main

func subarraySum(nums []int, k int) int {
	result := 0
	sum := 0
	sumMap := make(map[int]int)

	for _, value := range nums {
		sum += value
		if sum == k {
			result++
		} 

		prev, ok := sumMap[sum - k]
		if ok {
			result += prev
		}

		freq, ok := sumMap[sum]
		if ok {
			sumMap[sum] = freq + 1
		} else {
			sumMap[sum] = 1
		}
	}

  return result 
}
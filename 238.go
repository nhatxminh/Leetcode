package main


func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	total := 1
	zeroCount := 0
	zeroIndex := 0

	for i := range nums {
		if nums[i] == 0 {
			zeroCount++
			zeroIndex = i
		} else {
			total *= nums[i]
		}
	}
	
	if zeroCount > 1 {
		output = make([]int, len(nums))
	} else if zeroCount == 1 {
		output[zeroIndex] = total
	} else {
		for i := range nums {	
			output[i] = total/ nums[i]
		}
	}

	return output
}
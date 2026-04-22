package main


func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	blank := make([]int, len(nums))
	left := 1
	right := 1
	zeroCount := 0
	zeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
			zeroIndex = i
		} else {
			left *= nums[i]
			if i < len(nums) - 1 {
				output[i + 1] = left
			}
		}
	}
	
	if zeroCount > 1 {
		return blank
	} else if zeroCount == 1 {
		blank[zeroIndex] = left
		return blank
	}

	for i := len(nums) - 1; i > 0; i-- {
		right *= nums[i]
		if i == 1 {
			output[i - 1] = right
		} else {
			output[i - 1] *= right
		}
	}

	return output
}
package main 

func canJump(nums []int) bool {
	idx := -1
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums) - 1 {
			continue
		}
		if nums[i] == 0 {
			if idx == -1 {
				idx = i
			}
		} else {
			if nums[i] + i > idx {
				idx = -1
			}
		}
	}

	if idx == -1 {
		return true
	} else {
		return false 
	}
}
package main

import "math"

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for {
		mid := left + (right - left) / 2

		if nums[left] < nums[right] {
			return nums[left]
		}

		if left == right {
			return nums[left]
		}

		if mid > 0 && mid < len(nums)-1 && nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
			return nums[mid]
		}

		if right-left < 3 {
			return int(math.Min(math.Min(float64(nums[left]), float64(nums[right])), float64(nums[mid])))
		}

		if nums[left] > nums[right] && nums[mid] < nums[left] {
			right = mid - 1
		}

		if nums[left] > nums[right] && nums[mid] > nums[left] {
			left = mid + 1
		}
	}
}
package main

func getCommon(nums1 []int, nums2 []int) int {
	position1 := 0
	position2 := 0

	for position1 < len(nums1) || position2 < len(nums2) {
		if nums1[position1] == nums2[position2] {
			return nums1[position1]
		}

		if position1 < len(nums1) - 1 && nums1[position1] < nums2[position2] {
			position1++
		} else if position2 < len(nums2) - 1 && nums2[position2] < nums1[position1] {
			position2++
		} else {
			return -1
		}
	}

	return -1 
}
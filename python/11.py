from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        left = 0
        right = len(height) - 1
        maxValue = 0

        while left != right:
            maxValue = max((right - left) * min(height[left], height[right]), maxValue)

            if height[left] >= height[right]:
                right -= 1
            else: 
                left += 1
        return maxValue
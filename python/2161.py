from typing import List


class Solution:
    def pivotArray(self, nums: List[int], pivot: int) -> List[int]:
        left = []
        middle = []
        right = []

        for idx, num in enumerate(nums):
            if num == pivot:
                middle.append(num)
            elif num < pivot:
                left.append(num)
            elif num > pivot:
                right.append(num)
        
        return left + middle + right
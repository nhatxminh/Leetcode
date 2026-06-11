from typing import List


class Solution:
    def maxTotalValue(self, nums: List[int], k: int) -> int:
        min_value = min(nums)
        max_value = max(nums)
        return (max_value - min_value) * k
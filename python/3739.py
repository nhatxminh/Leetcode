import bisect
from typing import List


class FenwickTree:
    def __init__(self, length) -> None:
        self.length = length
        self.bit = [0] * (length + 1)
    
    def update(self, idx, value):
        while idx < self.length:
            self.bit[idx] += value
            idx += idx & -idx
    
    def query(self, idx) -> int:
        total = 0
        while idx > 0:
            total += self.bit[idx]
            idx -= idx & -idx
        return total

class Solution:
    def countMajoritySubarrays(self, nums: List[int], target: int) -> int:
        result = 0
        prefixes = [0 for i in range(len(nums) + 1)]

        for idx, num in enumerate(nums):
            prefixes[idx + 1] = prefixes[idx] + (1 if num == target else -1)

        sorted_prefix = sorted(list(set(prefixes)))
        fw = FenwickTree(len(sorted_prefix) + 1)

        for prefix in prefixes:
            idx = bisect.bisect_left(sorted_prefix, prefix) + 1
            result += fw.query(idx - 1)
            fw.update(idx, 1)
        return result
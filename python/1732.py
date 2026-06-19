from typing import List


class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        highest = 0
        curr = 0

        for _, move in enumerate(gain):
            curr += move
            highest = max(highest, curr)
        return highest
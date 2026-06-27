import math
from typing import List


class Solution:
    def maximumLength(self, nums: List[int]) -> int:
        result = 0
        freq: dict[int, int] = {}

        for num in nums:
            freq[num] = freq.get(num, 0) + 1

        for num in nums:
            total = 1
            curr = math.sqrt(num)
            while curr.is_integer() and curr != 1:
                curr = int(curr)
                if freq.get(curr, 0) >= 2:
                    total += 2
                    curr = math.sqrt(curr)
                else:
                    break 
            result = max(result, total)
        
        num_1 = freq.get(1, 0)
        count_1 = 1 + (num_1 - 1) // 2 * 2
        return max(result, count_1)
    
if __name__ == "__main__":
    solution = Solution()
    num = math.sqrt(5)
    print(num.is_integer())
    nums = [1,3,2,4]
    print(solution.maximumLength(nums))
from typing import List


class Solution:
    def leftRightDifference(self, nums: List[int]) -> List[int]:
        left = 0
        right = sum(nums)
        result = []

        for idx, curr in enumerate(nums):
            right -= curr
            result.append(abs(right - left))
            left += curr

        return result 
    
if __name__ == "__main__":
    solution = Solution()
    print(solution.leftRightDifference([10,4,8,3]))
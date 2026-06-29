from typing import List

class Solution:
    def maximumElementAfterDecrementingAndRearranging(self, arr: List[int]) -> int:
        length = len(arr)
        freq = [0] * (length + 1)
        max = 0

        for idx in range(length):
            freq[min(arr[idx], length)] += 1
        
        for idx, num in enumerate(freq):
            max = min(idx, max + num)
        return max
    
if __name__ == "__main__":
    solution = Solution()
    arr = [2,2,2,2,2,2,2,2,2,2,2,3,3,3,3,44,44,44,44,44,44,44,44,44,10000]
    print(solution.maximumElementAfterDecrementingAndRearranging(arr))
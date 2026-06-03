import sys
from typing import List


class Solution:
    def earliestFinishTime(self, landStartTime: List[int], landDuration: List[int], waterStartTime: List[int], waterDuration: List[int]) -> int:
        result = sys.maxsize
        minLandArrive = sys.maxsize
        minWaterArrive = sys.maxsize

        for idx, start in enumerate(landStartTime):
            minLandArrive = min(start + landDuration[idx], minLandArrive)

        for idx, start in enumerate(waterStartTime):
            minWaterArrive = min(start + waterDuration[idx], minWaterArrive)

        for idx, start in enumerate(landStartTime):
            result = min(max(minWaterArrive, start) + landDuration[idx], result)

        for idx, start in enumerate(waterStartTime):
            result = min(max(minLandArrive, start) + waterDuration[idx], result)
        return result
    
if __name__ == "__main__":
    solution = Solution()
    landStartTime = [2,8]
    landDuration = [4,1] 
    waterStartTime = [6] 
    waterDuration = [3]
    print(solution.earliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration))
import sys
from typing import List


class Solution:
    def earliestFinishTime(self, landStartTime: List[int], landDuration: List[int], waterStartTime: List[int], waterDuration: List[int]) -> int:
        result = sys.maxsize
        for idx_land, num_land in enumerate(landStartTime):
            for idx_water, num_water in enumerate(waterStartTime):
                smaller = min(max(num_water + waterDuration[idx_water], num_land) + landDuration[idx_land], max(num_land + landDuration[idx_land], num_water) + waterDuration[idx_water])
                result = min(smaller, result)
        return result
    
class Main: 
    solution = Solution()
    landStartTime = [2,8]
    landDuration = [4,1] 
    waterStartTime = [6] 
    waterDuration = [3]
    print(solution.earliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration))
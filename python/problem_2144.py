from typing import List


class Solution:
    def minimumCost(self, cost: List[int]) -> int:
        cost.sort()
        total = 0
        count = 0
        for i in range(len(cost) - 1, -1, -1):
            count += 1
            if count % 3 == 0:
                continue
            total += cost[i]
        return total

class Main: 
    cost = [6,5,7,9,2,2]
    solution = Solution()
    print(solution.minimumCost(cost=cost))
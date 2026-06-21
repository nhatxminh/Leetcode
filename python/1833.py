from typing import List


class Solution:
    def maxIceCream(self, costs: List[int], coins: int) -> int:
        highest = max(costs)
        count_arr = [0 for _ in range(highest + 1)]
        count = 0

        for _, cost in enumerate(costs):
            count_arr[cost] += 1
        
        for cost, freq in enumerate(count_arr):
            if cost == 0:
                continue

            can_buy = min(count_arr[cost], coins // cost)

            count += can_buy
            coins -= cost * freq

            if coins < 0:
                break
        return count
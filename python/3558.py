from collections import deque
from typing import List

MOD = 10 ** 9 + 7

class Solution:
    def assignEdgeWeights(self, edges: List[List[int]]) -> int:
        m: dict[int, List[int]] = {}
        bfs = deque()

        for _, edge in enumerate(edges):
            if edge[0] > edge[1]:
                start = edge[1]
                end = edge[0]
            else:
                start = edge[0]
                end = edge[1]
            if m.get(start) == None:
                m[start] = []
            m[start].append(end)
        
        bfs.append(1)
        curr_depth = 0
        curr_width = 0
        depth_map: dict[int, int] = {}
        depth_map[0] = 1
        while len(bfs) > 0:
            curr = bfs.popleft()
            next = m.get(curr)
            curr_width += 1

            if next != None:
                depth_map[curr_depth + 1] = depth_map.get(curr_depth + 1, 0) + len(next)
                bfs.extend(next)
            
            if curr_width == depth_map[curr_depth]:
                curr_depth += 1
                curr_width = 0 

        return int(2**(curr_depth - 2) % MOD)

if __name__ == "__main__":
    solution = Solution()
    edges = [[1,2],[1,3],[3,4],[3,5]]
    print(solution.assignEdgeWeights(edges))
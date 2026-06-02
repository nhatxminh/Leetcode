from collections import deque

class Solution:
    def canReach(self, s: str, minJump: int, maxJump: int) -> bool:
        dq: deque[int] = deque()
        dq.append(0)

        if s[-1] != '0':
            return False
        
        for i, char in enumerate(s):
            if len(dq) == 0:
                return False
            if dq[0] + maxJump == i:
                if char == '0':
                    dq.append(i)
                dq.popleft()
                continue
            if dq[0] + maxJump > i and dq[0] + minJump <= i and char == '0':
                dq.append(i)
    
        return dq[-1] == len(s) -1 

class Main:
    soltuion = Solution()
    s = "011111000111000001011111010"
    minJump = 6
    maxJump = 8
    print(soltuion.canReach(s, minJump, maxJump))
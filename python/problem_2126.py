from typing import List


class Solution:
    def asteroidsDestroyed(self, mass: int, asteroids: List[int]) -> bool:
        asteroids.sort()
        for asteroid in asteroids: 
            if mass >= asteroid: 
                mass += asteroid
            else:
                return False
        return True

class Main:
    mass = 84
    asteroids = [156,197,192,14,97,160,14,5]
    solution = Solution()
    print(solution.asteroidsDestroyed(mass=mass, asteroids=asteroids))
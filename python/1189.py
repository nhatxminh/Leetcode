class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        string = {"b", "a", "l", "o", "n"}
        m: dict[str, int] = {}

        for _, char in enumerate(text):
            if char in string:
                m[char] = m.get(char, 0) + 1

        b = m["b"]
        a = m["a"]
        l = m["l"] // 2
        o = m["o"] // 2
        n = m["n"]
        return min(b, a, l, o, n)
    
if __name__ == "__main__":
    solution = Solution()
    print(solution.maxNumberOfBalloons("nlaebolko"))
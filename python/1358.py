class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        result = 0
        freq = [-1] * 3

        for idx, char in enumerate(s):
            freq[ord(char) - 97] = idx
            a = freq[0]
            b = freq[1]
            c = freq[2]
            if a != -1 and b != -1 and c!= -1:
                result += min(a, b, c) + 1
        return result
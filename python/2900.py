from typing import List


class Solution:
    def getLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        result: list[str] = []
        prev = 0

        for idx, word in enumerate(words):
            if idx == 0:
                prev = groups[idx]
                result.append(word)
                continue
            if groups[idx] != prev:
                prev = groups[idx]
                result.append(word)
        return result 
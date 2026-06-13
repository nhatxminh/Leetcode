from typing import List

class Solution:
    def mapWordWeights(self, words: List[str], weights: List[int]) -> str:
        string = ""
        for _, word in enumerate(words):
            sum = 0
            for _, char in enumerate(word):
                sum += weights[ord(char) - ord("a")]
            string += chr(122 - (sum % 26))
        return string
    
if __name__ == "__main__":
    solution = Solution()
    words = ["abcd","def","xyz"] 
    weights = [5,3,12,14,1,2,3,2,10,6,6,9,7,8,7,10,8,9,6,9,9,8,3,7,7,2]
    print(solution.mapWordWeights(words, weights))
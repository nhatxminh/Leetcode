arr = []
peak = 0

for j in range(0, 130 + 1, 1):
    result = 0
    num_str = str(j)
    if j < 100:
        arr.append(0)
    else: 
        for i in range(0, len(num_str), 1):
            if i == 0 or i == len(num_str) - 1:
                continue
            if num_str[i] > num_str[i - 1] and num_str[i] > num_str[i + 1]:
                result += 1
            elif num_str[i] < num_str[i - 1] and num_str[i] < num_str[i + 1]:
                result += 1
        arr.append(result + arr[j - 1])


class Solution:
    def totalWaviness(self, num1: int, num2: int) -> int:
        return arr[num2] - arr[num1 - 1]
    

if __name__ == "__main__":
    print(arr[120])
    # solution = Solution()
    # print(solution.totalWaviness(120, 130))
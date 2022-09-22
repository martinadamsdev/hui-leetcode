from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        pos = 0
        temp = []
        result = [[]]
        backtrack(nums, pos, temp, result)
        return result


def backtrack(nums: List[int], pos: int, temp: List[int], result: List[List[int]]):
    ans = temp[pos:]
    result = result.append(ans)
    print(result)

    for i in range(len(nums)):
        temp = ans.append(nums[i])
        # backtrack(nums, i+1, temp, result)
        # temp = temp[0: len(temp) - 1]


if __name__ == '__main__':
    solution = Solution()
    solution.subsets([1, 2, 3])

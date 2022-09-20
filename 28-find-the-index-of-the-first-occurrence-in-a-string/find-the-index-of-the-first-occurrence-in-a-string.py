class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        index = -1
        for i in range(len(haystack) - len(needle) + 1):
            if haystack[i:i+len(needle)] == needle:
                index = i
                break
        return index


if __name__ == '__main__':
    solution = Solution()
    print(solution.strStr("sadbut", "sad"))
    print(solution.strStr("sabbutsad", "sad"))
    print(solution.strStr("leetcode", "leeto"))

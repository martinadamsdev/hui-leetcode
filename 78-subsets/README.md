

```go
func lengthOfLongestSubstring(s string) int {
	l, r, max := 0, 0, 0
	for i := range s {
		r = i + 1
		index := strings.Index(s[l:i], string(s[i]))
		if index != -1 {
			l += index + 1
		}
		if len(s[l:r]) > max {
			max = len(s[l:r])
		}
	}
	return max
}
```

```python3
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if not s:
            return 0
        left = 0
        lookup = set()
        n = len(s)
        max_len = 0
        cur_len = 0

        for i in range(n):
            cur_len += 1

            while s[i] in lookup:
                lookup.remove(s[left])
                left += 1
                cur_len -= 1

            if cur_len > max_len:
                max_len = cur_len

            lookup.add(s[i])

        return max_len
```
package main

import (
	"strings"
)

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

func main() {
	test := "abcabcbb"
	println(lengthOfLongestSubstring(test))
}

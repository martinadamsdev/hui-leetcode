package main

//func strStr(haystack string, needle string) int {
//	return strings.Index(haystack, needle)
//}

// 滑动窗口
func strStr(haystack string, needle string) int {
	l := -1
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			l = i
			break
		}
	}
	return l
}

func main() {
	print(strStr("sadbut", "sad"))
}

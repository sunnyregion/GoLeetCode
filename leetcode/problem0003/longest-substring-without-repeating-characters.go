package problem0003

import "strings"

func lengthOfLongestSubstring(s string) int {
	var max, start, end, length = 0, 0, 0, len(s)
	for start < length && end < length {
		if !strings.Contains(s[start:end], string(s[end])) {
			end = end + 1
			if end-start > max {
				max = end - start
			}
		} else {
			start = start + 1
		}
	}
	return max
}

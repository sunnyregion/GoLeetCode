package problem0005

func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i = i + 1 {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := len1
		if len2 > len1 {
			len = len2
		}
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L = L - 1
		R = R + 1
	}
	return R - L - 1
}

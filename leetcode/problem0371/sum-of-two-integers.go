package problem0371

func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	for a != 0 {
		a, b = (a&b)<<1, a^b
	}
	return b
}

package problem0371

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/sunnyregion/GoLeetCode/leetcode/problem0371"
)

// tcs is testcase slice
var tcs = []struct {
	a   int
	b   int
	ans int
}{

	{0, 2, 2},

	{1, 0, 1},

	{1, 2, 3},

	{10, 222, 232},
	// 可以有多个 testcase
}

func Test_getSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, getSum(tc.a, tc.b), "输入:%v", tc)
	}
}

func Benchmark_getSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			getSum(tc.a, tc.b)
		}
	}
}

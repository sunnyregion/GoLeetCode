package problem0006

import (
	"bytes"
)

func convert(s string, nRows int) string {
	if nRows == 1 || len(s) <= nRows {
		return s
	}

	res := bytes.Buffer{}
	// p pace 步距
	p := nRows*2 - 2

	// 处理第一行
	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	// 处理中间的行
	for n := 1; n <= nRows-2; n++ {
		// 添加r行的第一个字符
		res.WriteByte(s[n])

		for k := p; k-n < len(s); k += p {
			res.WriteByte(s[k-n])
			if k+n < len(s) {
				res.WriteByte(s[k+n])
			}
		}
	}

	// 处理最后一行
	for i := nRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	return res.String()
}

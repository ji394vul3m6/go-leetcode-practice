package solutions

import (
	"bytes"
	"fmt"
)

func ConvertZigZag() bool {
	type testCase struct {
		Input  string
		Rows   int
		Except string
	}
	testCases := []testCase{
		testCase{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	}
	ret := true
	for idx, testcase := range testCases {
		result := convertZigZag(testcase.Input, testcase.Rows)
		if result != testcase.Except {
			ret = false
			fmt.Printf("Case %d fail: %s, %d, %s, get '%s'\n", idx, testcase.Input, testcase.Rows, testcase.Except, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func convertZigZag(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	buf := make([]bytes.Buffer, numRows)
	currentRow := 0
	diff := 1
	for i := 0; i < len(s); i++ {
		buf[currentRow].WriteString(string(s[i]))
		if diff == 1 && currentRow == numRows-1 {
			diff = -1
		} else if diff == -1 && currentRow == 0 {
			diff = 1
		}
		currentRow += diff
	}

	for i := 1; i < len(buf); i++ {
		buf[0].WriteString(buf[i].String())
	}

	return buf[0].String()
}

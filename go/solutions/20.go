package solutions

import (
	"fmt"
)

func IsValidParentheses() bool {
	type testCase struct {
		Input  string
		Expect bool
	}
	testCases := []testCase{
		testCase{"{{[[(())]]}}", true},
		testCase{"]", false},
		testCase{"", true},
	}
	ret := true
	for idx, testcase := range testCases {
		result := isValidParentheses(testcase.Input)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %s, %t, get %t\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func isValidParentheses(s string) bool {
	pairs := map[byte]byte{
		'[': ']',
		']': '[',
		'(': ')',
		')': '(',
		'{': '}',
		'}': '{',
	}
	pushAction := map[byte]bool{
		'[': true,
		']': false,
		'(': true,
		')': false,
		'{': true,
		'}': false,
	}
	stack := make(map[int]byte)

	curIdx := 0
	for i := 0; i < len(s); i++ {
		if pushAction[s[i]] {
			stack[curIdx] = s[i]
			curIdx++
		} else {
			curIdx--
			if stack[curIdx] != pairs[s[i]] {
				return false
			}
		}
	}

	return curIdx == 0
}

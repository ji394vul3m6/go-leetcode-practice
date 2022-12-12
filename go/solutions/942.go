package solutions

import (
	"fmt"

	"go-leetcode-pratice/solutions/util"
)

func DiStringMatch() bool {
	type testCase struct {
		Input  string
		Expect []int
	}
	testCases := []testCase{
		testCase{"IDID", []int{0, 4, 1, 3, 2}},
		testCase{"III", []int{0, 1, 2, 3}},
		testCase{"DDI", []int{3, 2, 0, 1}},
	}
	ret := true
	for idx, testcase := range testCases {
		result := diStringMatch(testcase.Input)
		if !util.TestIntSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail: %s %#v, get %#v\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func diStringMatch(S string) []int {
	start := 0
	end := len(S)

	ret := make([]int, 0, len(S))
	for _, r := range S {
		switch r {
		case 'I':
			ret = append(ret, start)
			start++
		case 'D':
			ret = append(ret, end)
			end--
		}
	}
	ret = append(ret, start)
	return ret
}

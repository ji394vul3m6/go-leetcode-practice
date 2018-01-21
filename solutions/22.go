package solutions

import (
	"fmt"
	"litttlebear/leetcode/util"
)

func GenerateParenthesis() bool {
	type testCase struct {
		Input  int
		Except []string
	}
	testCases := []testCase{
		testCase{3, []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}},
		testCase{4, []string{
			"(((())))",
			"((()()))",
			"((())())",
			"((()))()",
			"(()(()))",
			"(()()())",
			"(()())()",
			"(())(())",
			"(())()()",
			"()((()))",
			"()(()())",
			"()(())()",
			"()()(())",
			"()()()()",
		}},
	}
	ret := true
	for idx, testcase := range testCases {
		result := generateParenthesis(testcase.Input)
		if !util.TestStringSliceEqual(result, testcase.Except) {
			ret = false
			fmt.Printf("Case %d fail: %d, %#v, get %#v\n", idx, testcase.Input, testcase.Except, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func generateParenthesis(n int) []string {
	type Running struct {
		Val   string
		Start int
		End   int
	}

	runnings := []Running{
		Running{"(", 1, 0},
	}

	for i := 1; i < 2*n; i++ {
		newRunnings := []Running{}
		for _, running := range runnings {
			if running.Start < n {
				newRunnings = append(newRunnings, Running{
					running.Val + "(", running.Start + 1, running.End,
				})
			}
			if running.End < running.Start {
				newRunnings = append(newRunnings, Running{
					running.Val + ")", running.Start, running.End + 1,
				})
			}
		}
		runnings = newRunnings
	}

	ret := []string{}
	for _, running := range runnings {
		ret = append(ret, running.Val)
	}

	return ret
}

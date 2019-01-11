package solutions

import (
	"fmt"
	"strconv"
)

func BaseballGame() bool {
	type testCase struct {
		Input  []string
		Expect int
	}
	testCases := []testCase{
		testCase{[]string{"5", "2", "C", "D", "+"}, 30},
		testCase{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
	}
	ret := true
	for idx, testcase := range testCases {
		result := calPoints(testcase.Input)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %#v, %d, get %d\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func calPoints(ops []string) int {
	ret := 0
	stack := []int{}
	for _, op := range ops {
		val := 0
		switch op {
		case "+":
			val = (stack[len(stack)-1] + stack[len(stack)-2])
			stack = append(stack, val)
		case "D":
			val = stack[len(stack)-1]
			val *= 2
			stack = append(stack, val)
		case "C":
			val = -1 * stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			var err error
			val, err = strconv.Atoi(op)
			if err != nil {
				continue
			}
			stack = append(stack, val)
		}
		ret += val
	}
	return ret
}

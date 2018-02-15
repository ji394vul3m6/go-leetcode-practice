package solutions

import (
	"fmt"
	"strings"
)

func Calculate() bool {
	type testCase struct {
		String string
		Expect int
	}
	testCases := []testCase{
		testCase{"3+2*2", 7},
		testCase{" 3 ", 3},
		testCase{" 3/2 ", 1},
		testCase{" 3+5 / 2 ", 5},
		testCase{" 3+ 5 * 3 / 2 ", 10},
		testCase{" 10 + 5 * 3 / 2 ", 17},
	}
	ret := true
	for idx, testcase := range testCases {
		result := calculate(testcase.String)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %s, %d, get %d\n\n", idx, testcase.String, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n\n", idx)
		}
	}
	return ret
}

func calculate(s string) int {
	pure := strings.Replace(s, " ", "", -1)
	type Operation struct {
		Op     byte
		Number int
	}
	operations := []Operation{}
	firstNum, next := getNextNum(pure, 0)
	operations = append(operations, Operation{
		'+',
		firstNum,
	})
	for next < len(pure) {
		op := pure[next]
		num := 0
		num, next = getNextNum(pure, next+1)
		operations = append(operations, Operation{
			op,
			num,
		})
	}

	idx := 0
	ret := 0
	for ; idx < len(operations); idx++ {
		op := operations[idx]
		positive := true
		temp := -1
		if op.Op == '+' || op.Op == '-' {
			// fmt.Printf("Start from %d\n", op.Number)
			positive = (op.Op == '+')
			temp = op.Number
			for ; idx < len(operations)-1; idx++ {
				nextOp := operations[idx+1]
				if nextOp.Op == '*' {
					// fmt.Printf("Multiply %d\n", nextOp.Number)
					temp *= nextOp.Number
				} else if nextOp.Op == '/' {
					// fmt.Printf("Divide %d\n", nextOp.Number)
					temp /= nextOp.Number
				} else {
					break
				}
			}
			if !positive {
				temp *= -1
			}
			// fmt.Printf("Partial value %d\n", temp)
			ret += temp
		} else {
			fmt.Println("Error format")
		}
	}

	return ret
}

func getNextNum(chars string, idx int) (int, int) {
	ret := 0
	i := idx
	for ; i < len(chars) && chars[i]-'0' >= 0 && chars[i]-'0' < 10; i++ {
		ret *= 10
		ret += int(chars[i] - '0')
	}
	return ret, i
}

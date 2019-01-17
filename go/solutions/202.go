package solutions

import "fmt"

func IsHappy() bool {
	type testCase struct {
		Input  int
		Expect bool
	}
	testCases := []testCase{
		testCase{19, true},
		testCase{2, false},
	}
	ret := true
	for idx, testcase := range testCases {
		result := isHappy(testcase.Input)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %d, %t, get %t\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

var (
	numMap = map[int]int{}
)

func isHappy(n int) bool {
	var a int
	happened := map[int]bool{}
	num := n
	for num != 1 {
		sum := 0
		happened[num] = true
		if _, ok := numMap[num]; ok {
			num = numMap[num]
		} else {
			for num != 0 {
				num, a = num/10, num%10
				sum += a * a
			}
		}
		numMap[num] = sum
		num = sum
		if happened[num] {
			return false
		}
	}
	return true
}

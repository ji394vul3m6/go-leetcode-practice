package solutions

import (
	"fmt"
)

func ClimbStairs() bool {
	type testCase struct {
		Stairs int
		Except int
	}
	testCases := []testCase{
		testCase{1, 1},
		testCase{2, 2},
		testCase{3, 3},
		testCase{4, 5},
	}
	ret := true
	for idx, testcase := range testCases {
		result := climbStairs(testcase.Stairs)
		if result != testcase.Except {
			ret = false
			fmt.Printf("Case %d fail: %d, %d, get %d\n", idx, testcase.Stairs, testcase.Except, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func climbStairs(n int) int {
	ret := make([]int, n)
	if n <= 2 {
		return n
	}
	ret[0] = 1
	ret[1] = 2
	for i := 2; i < n; i++ {
		ret[i] = ret[i-1] + ret[i-2]
	}
	return ret[n-1]
}

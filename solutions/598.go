package solutions

import "fmt"

func RangeAdditionII() bool {
	type testCase struct {
		M      int
		N      int
		OP     [][]int
		Expect int
	}
	testCases := []testCase{
		testCase{3, 3, [][]int{[]int{2, 2}, []int{3, 3}}, 4},
		testCase{3, 4, [][]int{[]int{2, 3}, []int{3, 2}}, 4},
	}
	ret := true
	for idx, testcase := range testCases {
		result := maxCount(testcase.M, testcase.N, testcase.OP)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: Input: %d, %d, %+v\nExcept: %d, get %d\n",
				idx, testcase.M, testcase.N, testcase.OP, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	minM, minN := ops[0][0], ops[0][1]
	if minM > m {
		minM = m
	}
	if minN > n {
		minN = n
	}
	for _, op := range ops {
		if minM > op[0] {
			minM = op[0]
		}
		if minN > op[1] {
			minN = op[1]
		}
	}
	return minM * minN
}

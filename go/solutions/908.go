package solutions

import (
	"fmt"
	"sort"
)

func SmallestRangeI() bool {
	type testCase struct {
		Nums   []int
		K      int
		Expect int
	}
	testCases := []testCase{
		testCase{[]int{1}, 0, 0},
		testCase{[]int{0, 10}, 2, 6},
		testCase{[]int{1, 3, 6}, 3, 0},
	}
	ret := true
	for idx, testcase := range testCases {
		result := smallestRangeI(testcase.Nums, testcase.K)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %#v, %d, %d, get %d\n", idx, testcase.Nums, testcase.K, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func smallestRangeI(A []int, K int) int {
	sort.Ints(A)
	min := A[0]
	max := A[len(A)-1]

	ret := (max - min) - (2 * K)
	if ret < 0 {
		ret = 0
	}

	return ret
}

package solutions

import (
	"fmt"
)

func CanJump() bool {
	type testCase struct {
		Nums   []int
		Expect bool
	}
	testCases := []testCase{
		testCase{[]int{2, 3, 1, 1, 4}, true},
		testCase{[]int{3, 2, 1, 0, 4}, false},
	}
	ret := true
	for idx, testcase := range testCases {
		result := canJump(testcase.Nums)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %#v, %t, get %t\n", idx, testcase.Nums, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}
func canJump(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}

	if length == 1 {
		return true
	}

	testIdx := length - 1
	for idx := length - 1; idx >= 0; idx-- {
		// fmt.Printf("Check %d, jump range: %d\n", idx, idx+nums[idx])
		if idx+nums[idx] >= testIdx {
			testIdx = idx
		}
	}
	return testIdx == 0
}

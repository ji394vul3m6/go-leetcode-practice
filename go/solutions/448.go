package solutions

import (
	"fmt"

	"./util"
)

func FindDisappearedNumbers() bool {
	type testCase struct {
		Input  []int
		Expect []int
	}
	testCases := []testCase{
		testCase{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
	}
	ret := true
	for idx, testcase := range testCases {
		result := findDisappearedNumbers(testcase.Input)
		if !util.TestIntSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail: %+v %+v, get %#v\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func findDisappearedNumbers(nums []int) []int {
	ret := make([]int, len(nums))
	for idx := range ret {
		t := nums[idx]
		for t != 0 && ret[t-1] != 1 {
			ret[t-1], t = 1, nums[t-1]
		}
	}
	filtered := ret[:0]
	for idx := range ret {
		if ret[idx] != 1 {
			filtered = append(filtered, idx+1)
		}
	}
	return filtered
}

package solutions

import (
	"fmt"
)

func RemoveDuplicates() bool {
	type testCase struct {
		Input  []int
		Expect int
	}
	testCases := []testCase{
		testCase{[]int{1, 1, 2, 2, 3}, 3},
	}
	ret := true
	for idx, testcase := range testCases {
		result := removeDuplicates(testcase.Input)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %#v, %d, get %d\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("%#v\n", testcase.Input)
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func removeDuplicates(nums []int) int {
	writeIdx := 0
	idx := 0
	for idx < len(nums)-1 {
		idx++
		if nums[idx] != nums[idx-1] {
			writeIdx++
			nums[writeIdx] = nums[idx]
		}
	}
	return writeIdx + 1
}

package solutions

import (
	"fmt"
)

func RemoveSpecificElement() bool {
	type testCase struct {
		Input    []int
		Specific int
		Except   int
	}
	testCases := []testCase{
		testCase{[]int{1, 1, 2, 2, 3}, 2, 3},
		testCase{[]int{1, 2, 3, 3, 4, 5}, 3, 4},
		testCase{[]int{1, 2, 3, 3, 5, 5}, 5, 4},
	}
	ret := true
	for idx, testcase := range testCases {
		result := removeSpecificElement(testcase.Input, testcase.Specific)
		if result != testcase.Except {
			ret = false
			fmt.Printf("Case %d fail: %#v, %d, get %d\n", idx, testcase.Input, testcase.Except, result)
		} else {
			// fmt.Printf("%#v\n", testcase.Input)
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func removeSpecificElement(nums []int, val int) int {
	writeIdx := 0
	idx := 0
	for idx <= len(nums)-1 {
		if nums[idx] != val {
			nums[writeIdx] = nums[idx]
			writeIdx++
		}
		idx++
	}
	return writeIdx
}

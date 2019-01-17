package solutions

import "fmt"

func FindLengthOfLCIS() bool {
	type testCase struct {
		Input  []int
		Expect int
	}
	testCases := []testCase{
		testCase{[]int{1, 3, 2, 4, 5, 3, 3}, 3},
		testCase{[]int{1, 3, 5, 4, 7}, 3},
		testCase{[]int{2, 2, 2, 2, 2}, 1},
		testCase{[]int{1, 2, 3, 4, 5}, 5},
	}
	ret := true
	for idx, testcase := range testCases {
		result := findLengthOfLCIS(testcase.Input)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %#v, %d, get %d\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func findLengthOfLCIS(nums []int) int {
	max := 0
	start := 0
	for idx := range nums {
		if idx == len(nums)-1 {
			if idx-start+1 > max {
				max = idx - start + 1
			}
			continue
		}
		if nums[idx+1] <= nums[idx] {
			if idx-start+1 > max {
				max = idx - start + 1
			}
			start = idx + 1
		}
	}
	return max
}

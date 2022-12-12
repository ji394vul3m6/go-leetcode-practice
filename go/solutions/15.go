package solutions

import (
	"fmt"
	"go-leetcode-pratice/solutions/util"
	"sort"
)

func TreeSum() bool {
	type testCase struct {
		Nums   []int
		Expect [][]int
	}
	testCases := []testCase{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
	}
	ret := true
	for idx, testcase := range testCases {
		result := threeSum(testcase.Nums)
		if !util.TestDoubleIntSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail: %#v, %#v, get %#v\n", idx, testcase.Nums, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	for i := range nums {
		if i+1 >= len(nums) {
			break
		}
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; k > j; {
			if j > i+1 && nums[j] == nums[j-1] {
				j += 1
				continue
			}
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k -= 1
				continue
			}

			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j += 1
				k -= 1
			} else if sum > 0 {
				k -= 1
			} else if sum < 0 {
				j += 1
			}
		}
	}
	return ret
}

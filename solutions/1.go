package solutions

import (
	"fmt"
	"litttlebear/leetcode/util"
)

func TwoSum() bool {
	type testCase struct {
		Nums   []int
		Target int
		Expect []int
	}
	testCases := []testCase{
		testCase{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		testCase{[]int{5, 4, 6}, 10, []int{1, 2}},
	}
	ret := true
	for idx, testcase := range testCases {
		result := twoSum(testcase.Nums, testcase.Target)
		if !util.TestIntSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail: %#v, %#v, %#v, get %#v\n", idx, testcase.Nums, testcase.Target, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func twoSum(nums []int, target int) []int {
	temp := make(map[int][]int)
	for idx, n := range nums {
		if temp[n] == nil {
			temp[n] = []int{}
		}
		temp[n] = append(temp[n], idx)
	}

	for idx, n := range nums {
		remain := target - n
		if remain == n {
			if len(temp[n]) >= 2 {
				return []int{temp[n][0], temp[n][1]}
			}
		} else if idxes, ok := temp[remain]; ok {
			return []int{idx, idxes[0]}
		}
	}

	return []int{}
}

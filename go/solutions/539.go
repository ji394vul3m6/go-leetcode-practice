package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func MinimumTimeDifference() bool {
	type testCase struct {
		TimePoints []string
		Expect     int
	}

	testCases := []testCase{
		{[]string{"11:00", "13:00", "16:00"}, 120},
		{[]string{"23:59", "00:00"}, 1},
		{[]string{"00:00", "23:59", "00:00"}, 0},
	}
	ret := true
	for idx, testcase := range testCases {
		result := findMinDifference(testcase.TimePoints)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: Input: %s\nExcept: %d, get %d\n", idx, testcase.TimePoints, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func findMinDifference(timePoints []string) int {
	nums := []int{}
	for _, point := range timePoints {
		parts := strings.Split(point, ":")
		if len(parts) != 2 {
			return -1
		}
		hour, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("%+v", err)
			return -1
		}
		min, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("%+v", err)
			return -1
		}
		nums = append(nums, hour*60+min)
	}
	sort.Ints(nums)
	diffs := []int{}
	for idx := range nums {
		if idx == len(nums)-1 {
			diffs = append(diffs, nums[0]+1440-nums[idx])
		} else {
			diffs = append(diffs, nums[idx+1]-nums[idx])
		}
	}
	sort.Ints(diffs)
	return diffs[0]
}

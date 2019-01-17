package solutions

import (
	"fmt"
	"sort"
)

func FindMinInRotateSorted() bool {
	type testCase struct {
		Input  []int
		Expect int
	}
	testCases := []testCase{
		testCase{[]int{1}, 1},
		testCase{[]int{1, 2}, 1},
		testCase{[]int{3, 4, 5, 1, 2}, 1},
		testCase{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	}
	ret := true
	for idx, testcase := range testCases {
		result := findMin3(testcase.Input)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %#v, %d, get %d\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			// fmt.Printf("%#v\n", testcase.Input)
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func findMin1(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

func findMin2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	mid := len(nums) / 2
	if nums[mid] < nums[mid-1] {
		return nums[mid]
	}

	if nums[mid-1] < nums[len(nums)-1] {
		return findMin2(nums[0:mid])
	} else {
		return findMin2(nums[mid:])
	}
}

func findMin3(nums []int) int {
	var findMin func(start, end int) int
	findMin = func(start, end int) int {
		if end-start == 1 {
			return nums[start]
		}
		mid := (end + start) / 2
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[mid-1] < nums[len(nums)-1] {
			return findMin(start, mid)
		}
		return findMin(mid, end)
	}
	return findMin(0, len(nums))
}

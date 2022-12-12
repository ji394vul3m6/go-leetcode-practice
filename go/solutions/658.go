package solutions

import (
	"fmt"

	"go-leetcode-pratice/solutions/util"
)

func FindClosestElements() bool {
	type testCase struct {
		Nums   []int
		K      int
		X      int
		Expect []int
	}
	testCases := []testCase{
		testCase{[]int{1}, 1, 0, []int{1}},
		testCase{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		testCase{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
		testCase{[]int{1, 2, 2, 2, 5, 5, 5, 8, 9, 9}, 4, 0, []int{1, 2, 2, 2}},
	}
	ret := true
	for idx, testcase := range testCases {
		result := findClosestElements(testcase.Nums, testcase.K, testcase.X)
		if !util.TestIntSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail: %#v, %d, %d\n\tExcept:%#v\n\tGet:%#v\n",
				idx, testcase.Nums, testcase.K, testcase.X, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func findClosestElements(arr []int, k int, x int) []int {
	start := getStartPos(arr, x)
	front := start
	back := start
	for back-front < k-1 {
		if front-1 >= 0 && back+1 < len(arr) {
			diff1 := x - arr[front-1]
			diff2 := arr[back+1] - x
			if diff1 <= diff2 {
				front--
			} else {
				back++
			}
		} else if front-1 >= 0 {
			front--
		} else if back+1 < len(arr) {
			back++
		} else {
			fmt.Println("k < len(arr)!")
			break
		}
	}
	return arr[front : back+1]
}

func findClosestElements2(arr []int, k int, x int) []int {
	start := getStartPos(arr, x)
	front := start
	back := start
	for back-front < k-1 {
		if front-1 >= 0 && back+1 < len(arr) {
			diff1 := x - arr[front-1]
			diff2 := arr[back+1] - x
			if diff1 <= diff2 {
				front--
			} else {
				back++
			}
		} else if front-1 >= 0 {
			front--
		} else if back+1 < len(arr) {
			back++
		} else {
			fmt.Println("k < len(arr)!")
			break
		}
	}
	return arr[front : back+1]
}

func getStartPos(arr []int, x int) int {
	start := 0
	end := len(arr)
	ret := 0

	for true {
		// fmt.Printf("check: %d ~ %d\n", start, end)
		idx := (start + end) / 2
		if idx == start {
			if arr[idx] > x {
				ret = idx - 1
			} else {
				ret = idx
			}
			break
		}

		if arr[idx] > x {
			end = idx
		} else if arr[idx] < x {
			start = idx
		} else {
			ret = idx
			break
		}
	}
	if ret < 0 {
		ret = 0
	}
	return ret
}

/*
Given a sorted array, two integers `k` and `x`, find the `k` closest elements to `x` in the array. The result should also be sorted in ascending order. If there is a tie, the smaller elements are always preferred.

Example 1:

Input: [1,2,3,4,5], k=4, x=3
Output: [1,2,3,4]

Example 2:

Input: [1,2,3,4,5], k=4, x=-1
Output: [1,2,3,4]

Note:

1.  The value k is positive and will always be smaller than the length of the sorted array.
2.  Length of the given array is positive and will not exceed 10^4
3.  Absolute value of elements in the array and x will not exceed 10^4
*/

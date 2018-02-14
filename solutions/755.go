package solutions

import (
	"fmt"
	"litttlebear/leetcode/util"
)

func PourWater() bool {
	type testCase struct {
		Heights []int
		Drops   int
		DropAt  int
		Expect  []int
	}
	testCases := []testCase{
		testCase{[]int{0}, 10, 0, []int{10}},
		testCase{[]int{2, 1, 1, 2, 1, 2, 2}, 4, 3, []int{2, 2, 2, 3, 2, 2, 2}},
		testCase{[]int{1, 2, 3, 4}, 2, 2, []int{2, 3, 3, 4}},
		testCase{[]int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1}, 2, 2, []int{2, 3, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1}},
		testCase{[]int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1}, 10, 2, []int{4, 4, 4, 4, 3, 3, 3, 3, 3, 4, 3, 2, 1}},
		testCase{[]int{14, 10, 10, 3, 13, 1, 2, 1, 2, 5}, 1, 0, []int{14, 10, 10, 4, 13, 1, 2, 1, 2, 5}},
		testCase{[]int{14, 9, 10, 9, 7, 9, 7, 5, 3, 2}, 7, 9, []int{14, 9, 10, 9, 7, 9, 7, 5, 6, 6}},
	}
	ret := true
	for idx, testcase := range testCases {
		result := pourWater(testcase.Heights, testcase.Drops, testcase.DropAt)
		if !util.TestIntSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail: %#v, %d, %d\n\texpect: %#v\n\tget: %#v\n",
				idx, testcase.Heights, testcase.Drops, testcase.DropAt, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func pourWater(heights []int, V int, K int) []int {
	remains := V
	ret := append([]int{}, heights...)

	for remains > 0 {
		dripWater(ret, K)
		remains--
	}

	return ret
}

func dripWater(heights []int, K int) {
	for i := K - 1; i >= 0 && heights[i] <= heights[K]; i-- {
		if (i == 0 && heights[i] < heights[K]) || (i > 0 && heights[i-1] > heights[i]) {
			j := i
			fmt.Printf("%d\n", i)
			for j <= K && heights[j] == heights[i] {
				j++
			}
			if j-1 != K {
				heights[j-1]++
				return
			}
		}
	}
	for i := K + 1; i < len(heights) && heights[i] <= heights[K]; i++ {
		if (i == len(heights)-1 && heights[i] < heights[K]) || (i < len(heights)-1 && heights[i+1] > heights[i]) {
			j := i
			for j >= K && heights[j] == heights[i] {
				j--
			}
			if j+1 != K {
				heights[j+1]++
				return
			}
		}
	}
	// fmt.Println("Drip on K")
	heights[K]++
}

func debugPrint(ret []int, pos int) {
	fmt.Printf("Debug:\n")
	temp := make([]string, len(ret))
	for idx := range ret {
		if idx == pos {
			temp[idx] = "v"
		} else {
			temp[idx] = " "
		}
	}

	for _, v := range temp {
		fmt.Printf("%s ", v)
	}
	fmt.Print("\n")
	for _, v := range ret {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}

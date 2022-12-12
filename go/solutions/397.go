package solutions

import (
	"fmt"

	"go-leetcode-pratice/solutions/util"
)

func IntegerReplacement() bool {
	type testCase struct {
		N      int
		Expect int
	}
	testCases := []testCase{
		testCase{8, 3},
		testCase{7, 4},
		testCase{57, 8},
	}
	ret := true
	for idx, testcase := range testCases {
		result := integerReplacement2(testcase.N)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %d, %d, get %d\n", idx, testcase.N, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func integerReplacement(n int) int {
	c := n
	extra := 0
	if c%2 == 0 {
		extra++
		c = c / 2
	}

	buffer := make([]int, c+2)
	buffer[1] = 0
	buffer[2] = 1

	for idx := 4; idx < c+2; idx += 2 {
		buffer[idx] = 1 + buffer[idx/2]

		opt1 := 1 + buffer[idx]
		opt2 := 1 + buffer[idx-2]
		buffer[idx-1] = util.Min([]int{opt1, opt2})
	}

	return buffer[c] + extra
}

func integerReplacement2(n int) int {
	candidates := []int{n}
	showed := map[int]bool{}

	showed[n] = true
	ret := 0
	for true {
		// fmt.Printf("%#v\n", candidates)
		next := []int{}
		finish := false
		for _, number := range candidates {
			if number == 1 {
				finish = true
				break
			}
			if number%2 == 0 && !showed[number/2] {
				next = append(next, number/2)
				showed[number/2] = true
			} else {
				if !showed[number+1] {
					next = append(next, number+1)
					showed[number+1] = true
				}
				if !showed[number-1] {
					next = append(next, number-1)
					showed[number-1] = true
				}
			}
		}
		if finish {
			break
		}
		ret++
		candidates = next
	}
	return ret
}

/*
Given a positive integer *n* and you can do operations as follow:

1.  If *n* is even, replace *n* with `*n* / 2`.
2.  If *n* is odd, you can replace *n* with either `*n* + 1` or `*n* - 1`.

What is the minimum number of replacements needed for *n* to become 1?

Example 1:
Input:
8
Output:
3
Explanation:
8 -> 4 -> 2 -> 1

Example 2:
Input:
7
Output:
4
Explanation:
7 -> 8 -> 4 -> 2 -> 1
or
7 -> 6 -> 3 -> 2 -> 1
*/

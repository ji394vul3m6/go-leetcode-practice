package solutions

import (
	"fmt"
	"math"
)

func UniquePaths() bool {
	type testCase struct {
		Width  int
		Height int
		Except int
	}
	testCases := []testCase{
		testCase{1, 1, 1},
		testCase{2, 2, 2},
		testCase{3, 3, 6},
		testCase{5, 10, 715},
		testCase{59, 5, 557845},
		// when overflow, my function will return 0
		testCase{100, 50, 0},
	}
	ret := true
	for idx, testcase := range testCases {
		result1 := uniquePaths(testcase.Width, testcase.Height)
		result2 := uniquePaths2(testcase.Width, testcase.Height)
		if result1 != testcase.Except || result2 != testcase.Except {
			ret = false
			fmt.Printf("Case %d fail: %d, %d, %d, get %d, %d\n", idx, testcase.Width, testcase.Height, testcase.Except, result1, result2)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}

	calc := make([][]int64, m)
	for idx := range calc {
		calc[idx] = make([]int64, n)
	}

	for idx := range calc[0] {
		calc[0][idx] = 1
	}

	for i := 1; i < len(calc); i++ {
		for j := range calc[i] {
			if j == 0 {
				calc[i][j] = 1
			} else {
				calc[i][j] = calc[i-1][j] + calc[i][j-1]
			}
			if calc[i][j] < 0 || calc[i][j] > math.MaxInt32 {
				// overflow in this functino prototype
				fmt.Printf("Overflow when calculating at %d %d\n", i, j)
				return 0
			}
		}
	}
	return int(calc[m-1][n-1])
}

func uniquePaths2(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}

	var ret int64 = 1
	n1, n2 := int64(m-1), int64(n-1)
	if n1 < n2 {
		n1, n2 = n2, n1
	}

	i := int64(n1) + int64(n2)
	var j int64 = 2
	for ; i > n1; i-- {
		// fmt.Printf(" * %d \n", i)
		ret *= i
		for ret%j == 0 && j <= n2 {
			// fmt.Printf(" / %d \n", j)
			ret /= j
			j++
		}
	}
	if ret > math.MaxInt32 {
		// overflow in this functino prototype
		return 0
	}
	return int(ret)
}

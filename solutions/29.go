package solutions

import "fmt"

func SelfDivide() bool {
	type testCase struct {
		Dividend int
		Divisor  int
		Except   int
	}
	testCases := []testCase{
		testCase{17, 2, 8},
		testCase{17, 1, 17},
		testCase{17, 3, 5},
		testCase{17, -3, -5},
		testCase{-17, -3, 5},
		testCase{-17, 3, -5},
		testCase{-2147483648, -1, 2147483647},
	}
	ret := true
	for idx, testcase := range testCases {
		result := divide(testcase.Dividend, testcase.Divisor)
		if result != testcase.Except {
			ret = false
			fmt.Printf("Case %d fail: %d, %d, %d, get %d\n", idx, testcase.Dividend, testcase.Divisor, testcase.Except, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func divide(dividend int, divisor int) int {
	ret := 0
	isMinus := false
	ratio := 1
	m := dividend
	if m < 0 {
		m = m * -1
		isMinus = !isMinus
	}
	n := divisor
	if n < 0 {
		n = n * -1
		isMinus = !isMinus
	}

	if n == 1 {
		ret = m
	} else {
		for m-n >= 0 {
			m -= n
			ret += ratio

			if (n<<1)>>1 == n && m > (n<<1) {
				n = n << 1
				ratio = ratio << 1
			}
			for m < n && ratio > 1 {
				n = n >> 1
				ratio = ratio >> 1
			}
		}
	}

	if isMinus {
		ret = ret * -1
	}
	if ret > 2147483647 {
		return 2147483647
	}
	if ret < -2147483648 {
		return -2147483648
	}
	return ret
}

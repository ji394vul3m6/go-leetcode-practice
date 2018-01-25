package solutions

import (
	"fmt"
)

func IsPalindrome() bool {
	type testCase struct {
		Input  int
		Expect bool
	}
	testCases := []testCase{
		testCase{1234321, true},
		testCase{123321, true},
		testCase{1222321, false},
		testCase{442, false},
		testCase{-2147447412, false},
	}
	ret := true
	for idx, testcase := range testCases {
		result := isPalindrome(testcase.Input)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %d, %t, get %t\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digitArray := []int{}

	tmp := x
	for tmp != 0 {
		digit := tmp % 10
		digitArray = append(digitArray, digit)
		tmp = tmp / 10
	}

	l := len(digitArray)
	for i := 0; i < l/2; i++ {
		if digitArray[i] != digitArray[l-1-i] {
			return false
		}
	}

	return true
}

package solutions

import (
	"fmt"
	"strconv"

	"./util"
)

func FizzBuzz() bool {
	type testCase struct {
		Num    int
		Expect []string
	}
	testCases := []testCase{
		testCase{15, []string{
			"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
		}},
	}
	ret := true
	for idx, testcase := range testCases {
		result := fizzBuzz(testcase.Num)
		if !util.TestStringSliceEqual(testcase.Expect, result) {
			ret = false
			fmt.Printf("Case %d fail: %d, Except:\n%+v\nGet\n%+v\n", idx, testcase.Num, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

var (
	max    = 0
	valStr = map[int]string{}
)

func fizzBuzz(n int) []string {
	if n > max {
		for i := (max - max%3 + 3); i <= n; i += 3 {
			valStr[i] = "Fizz"
		}
		for i := (max - max%5 + 5); i <= n; i += 5 {
			if _, ok := valStr[i]; ok {
				valStr[i] = "FizzBuzz"
			} else {
				valStr[i] = "Buzz"
			}
		}
		max = n
	}
	ret := make([]string, n)
	for idx := range ret {
		v := valStr[idx+1]
		if v == "" {
			ret[idx] = strconv.Itoa(idx + 1)
		} else {
			ret[idx] = v
		}
	}
	return ret
}

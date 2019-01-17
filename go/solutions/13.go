package solutions

import (
	"fmt"
	"strings"
)

func RomanToInt() bool {
	type testCase struct {
		Input  string
		Expect int
	}
	testCases := []testCase{
		testCase{"III", 3},
		testCase{"IV", 4},
		testCase{"IX", 9},
		testCase{"LVIII", 58},
	}
	ret := true
	for idx, testcase := range testCases {
		result := romanToInt(testcase.Input)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %s, %d, get %d\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func romanToInt(s string) int {
	type replaceCase struct {
		Target string
		Rune   string
		Times  int
	}

	cases := []replaceCase{
		replaceCase{"IV", "I", 4},
		replaceCase{"IX", "I", 9},
		replaceCase{"XL", "X", 4},
		replaceCase{"XC", "X", 9},
		replaceCase{"CD", "C", 4},
		replaceCase{"CM", "C", 9},
	}

	valStr := s
	for _, c := range cases {
		valStr = strings.Replace(valStr, c.Target, strings.Repeat(c.Rune, c.Times), -1)
	}

	runeVal := map[rune]int{
		'I': 1,
		'X': 10,
		'C': 100,
		'M': 1000,
		'V': 5,
		'L': 50,
		'D': 500,
	}

	ret := 0
	for _, c := range valStr {
		ret += runeVal[c]
	}

	return ret
}

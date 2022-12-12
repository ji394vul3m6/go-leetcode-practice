package solutions

import (
	"fmt"

	"go-leetcode-pratice/solutions/util"
)

func LetterCombinations() bool {
	type testCase struct {
		Digits string
		Expect []string
	}
	testCases := []testCase{
		testCase{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	ret := true
	for idx, testcase := range testCases {
		result := letterCombinations(testcase.Digits)
		if !util.TestStringSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail: %#v, %#v, get %#v\n", idx, testcase.Digits, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func letterCombinations(digits string) []string {
	letters := map[byte][]string{
		'0': []string{" "},
		'1': []string{" "},
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	var ret []string
	for _, digit := range digits {
		choices := letters[byte(digit)]
		nextRound := []string{}
		if len(ret) == 0 {
			nextRound = append(nextRound, choices...)
		} else {
			for _, answer := range ret {
				for _, choice := range choices {
					nextRound = append(nextRound, answer+choice)
				}
			}
		}
		ret = nextRound
	}
	return ret
}

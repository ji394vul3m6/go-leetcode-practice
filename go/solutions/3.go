package solutions

import (
	"fmt"
)

func LengthOfLongestSubstring() bool {
	type testCase struct {
		String string
		Expect int
	}
	testCases := []testCase{
		testCase{"abcabcbb", 3},
		testCase{"bbbbb", 1},
		testCase{"pwwkew", 3},
	}
	ret := true
	for idx, testcase := range testCases {
		result := lengthOfLongestSubstring(testcase.String)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %s, %d, get %d\n", idx, testcase.String, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func lengthOfLongestSubstring(s string) int {
	contain := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if _, ok := contain[s[i]]; !ok {
			contain[s[i]] = false
		}
	}

	start := 0
	end := 0
	maxLen := 0

	for end < len(s) {
		if !contain[s[end]] {
			contain[s[end]] = true
			end++
		} else {
			contain[s[start]] = false
			start++
		}
		if end-start > maxLen {
			maxLen = end - start
		}
	}

	return maxLen
}

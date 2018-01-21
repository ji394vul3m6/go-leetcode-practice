package solutions

import (
	"fmt"
)

func LongestPalindrome() bool {
	type testCase struct {
		String string
		Except string
	}
	testCases := []testCase{
		testCase{"babad", "bab"},
		testCase{"cbbd", "bb"},
		testCase{"ccc", "ccc"},
		testCase{"aaaa", "aaaa"},
	}
	ret := true
	for idx, testcase := range testCases {
		result := longestPalindrome(testcase.String)
		if result != testcase.Except {
			ret = false
			fmt.Printf("Case %d fail: %s, '%s', get '%s'\n", idx, testcase.String, testcase.Except, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		partLongest := getLogestOddPalindromeFrom(s, i)
		if len(partLongest) > len(longest) {
			longest = partLongest
		}
		partLongest = getLogestEvenPalindromeFrom(s, i)
		if len(partLongest) > len(longest) {
			longest = partLongest
		}
	}
	return longest
}

func getLogestEvenPalindromeFrom(s string, start int) string {
	if len(s) < 2 {
		return ""
	}
	if start+1 >= len(s) {
		return ""
	}
	if s[start] != s[start+1] {
		return ""
	}

	shiftLen := 0
	for start-shiftLen >= 0 && start+shiftLen+1 < len(s) {
		if s[start-shiftLen] != s[start+shiftLen+1] {
			shiftLen--
			break
		}

		if start-shiftLen-1 >= 0 && start+shiftLen+2 < len(s) {
			shiftLen++
		} else {
			break
		}
	}

	return s[start-shiftLen : start+shiftLen+2]
}

func getLogestOddPalindromeFrom(s string, start int) string {
	shiftLen := 0
	for start-shiftLen >= 0 && start+shiftLen < len(s) {
		if s[start-shiftLen] != s[start+shiftLen] {
			shiftLen--
			break
		}
		if start-shiftLen-1 >= 0 && start+shiftLen+1 < len(s) {
			shiftLen++
		} else {
			break
		}
	}

	return s[start-shiftLen : start+shiftLen+1]
}

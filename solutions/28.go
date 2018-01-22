package solutions

import (
	"fmt"
)

func ImplementStrStr() bool {
	type testCase struct {
		Haystack string
		Needle   string
		Except   int
	}
	testCases := []testCase{
		testCase{"hello", "ll", 2},
		testCase{"aaaaa", "bba", -1},
		testCase{"", "", 0},
		testCase{"a", "a", 0},
	}
	ret := true
	for idx, testcase := range testCases {
		result := implementStrStr(testcase.Haystack, testcase.Needle)
		if result != testcase.Except {
			ret = false
			fmt.Printf("Case %d fail: %s, %s, %d, get %d\n", idx, testcase.Haystack, testcase.Needle, testcase.Except, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func implementStrStr(haystack string, needle string) int {
	lenN := len(needle)
	for i := 0; i <= len(haystack)-lenN; i++ {
		subStr := haystack[i : i+len(needle)]
		if subStr == needle {
			return i
		}
	}
	return -1
}

package solutions

import "fmt"

func CountBinarySubstrings() bool {
	type testCase struct {
		String string
		Expect int
	}
	testCases := []testCase{
		testCase{"00110011", 6},
		testCase{"10101", 4},
		testCase{"1010001", 4},
	}
	ret := true
	for idx, testcase := range testCases {
		result := countBinarySubstrings(testcase.String)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %s, %d, get %d\n", idx, testcase.String, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func countBinarySubstrings(s string) int {
	count := 0
	start := 0
	current := 1
	for current < len(s) {
		if s[current] == s[start] {
			current++
		} else {
			length := 1
			for current+length < len(s) && length < current-start {
				if s[current+length] == s[current] {
					length++
				} else {
					break
				}
			}
			count += length
			start = current
			current = current + length
		}
	}
	return count
}

package solutions

import (
	"fmt"
	"math"
	"sort"
)

var (
	sortedExisted = map[string]bool{}
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func init() {
	i := int64(1)
	for i < int64(math.MaxInt32) {
		s := SortString(fmt.Sprintf("%d", i))
		sortedExisted[s] = true
		i *= 2
	}
}

func ReorderedPowerOf2() bool {
	type testCase struct {
		Input  int
		Expect bool
	}
	testCases := []testCase{
		testCase{1, true},
		testCase{10, false},
		testCase{16, true},
		testCase{24, false},
		testCase{46, true},
	}
	ret := true
	for idx, testcase := range testCases {
		result := reorderedPowerOf2(testcase.Input)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %d, %t, get %t\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func reorderedPowerOf2(N int) bool {
	s := SortString(fmt.Sprintf("%d", N))
	_, ok := sortedExisted[s]
	return ok
}

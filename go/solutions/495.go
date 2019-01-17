package solutions

import (
	"fmt"
)

func FindPoisonedDuration() bool {
	type testCase struct {
		TimeSeries []int
		Duration   int
		Expect     int
	}
	testCases := []testCase{
		testCase{[]int{1, 4}, 2, 4},
		testCase{[]int{1, 2}, 2, 3},
	}
	ret := true
	for idx, testcase := range testCases {
		result := findPoisonedDuration(testcase.TimeSeries, testcase.Duration)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %#v, %d, %d, get %d\n", idx, testcase.TimeSeries, testcase.Duration, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	ret := 0

	for idx := range timeSeries {
		if idx == 0 {
			continue
		}
		plus := timeSeries[idx] - timeSeries[idx-1]
		if plus > duration {
			plus = duration
		}
		ret += plus
	}

	if len(timeSeries) > 0 {
		ret += duration
	}

	return ret
}

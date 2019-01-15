package solutions

import "fmt"

func SuperPow() bool {
	type testCase struct {
		N      int
		Power  []int
		Expect int
	}
	testCases := []testCase{
		testCase{2, []int{3}, 8},
		testCase{2, []int{1, 0}, 1024},
	}
	ret := true
	for idx, testcase := range testCases {
		result := superPow(testcase.N, testcase.Power)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %d, %+v, %d, get %d\n", idx, testcase.N, testcase.Power, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func superPow(a int, b []int) int {
	ret := 1
	base := a
	for idx := range b {
		temp := 1
		for i := 1; i <= 10; i++ {
			temp = (temp * base) % 1337
			if i == b[len(b)-idx-1] {
				ret = (ret * temp) % 1337
			}
		}
		base = temp
	}
	return ret
}

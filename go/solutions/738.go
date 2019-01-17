package solutions

import "fmt"

func MonotoneIncreasingDigits() bool {
	type testCase struct {
		Input  int
		Expect int
	}
	testCases := []testCase{
		testCase{10, 9},
		testCase{1234, 1234},
		testCase{332, 299},
		testCase{100, 99},
	}
	ret := true
	for idx, testcase := range testCases {
		result := monotoneIncreasingDigits(testcase.Input)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %d, %d, get %d\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func monotoneIncreasingDigits(N int) int {
	nums := []int{}
	number := N
	for number > 0 {
		t := number % 10
		nums = append([]int{t}, nums...)
		number = number / 10
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if nums[j] > nums[i] {
				for k := i; k < len(nums); k++ {
					nums[k] = 9
				}
				nums[i-1]--
				break
			}
		}
	}

	ret := 0
	for _, val := range nums {
		ret *= 10
		ret += val
	}

	return ret
}

package solutions

import "fmt"

func LemonadeChange() bool {
	type testCase struct {
		Nums   []int
		Expect bool
	}
	testCases := []testCase{
		testCase{[]int{5, 5, 5, 10, 20}, true},
		testCase{[]int{5, 5, 10}, true},
		testCase{[]int{10, 10}, false},
		testCase{[]int{5, 5, 10, 10, 20}, false},
		testCase{[]int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}, true},
	}
	ret := true
	for idx, testcase := range testCases {
		result := lemonadeChange(testcase.Nums)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %#v, %t, get %t\n", idx, testcase.Nums, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func lemonadeChange(bills []int) bool {
	coinCount := map[int]int{
		5:  0,
		10: 0,
		20: 0,
	}

	for _, val := range bills {
		coinCount[val]++
		switch val {
		case 10:
			if coinCount[5] <= 0 {
				return false
			}
			coinCount[5]--
		case 20:
			if coinCount[5] > 0 && coinCount[10] > 0 {
				coinCount[5]--
				coinCount[10]--
			} else if coinCount[5] >= 3 {
				coinCount[5] -= 3
			} else {
				return false
			}
		}
	}

	return true
}

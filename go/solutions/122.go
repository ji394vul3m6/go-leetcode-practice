package solutions

import (
	"fmt"
)

func MaxProfit() bool {
	type testCase struct {
		Prices []int
		Expect int
	}
	testCases := []testCase{
		testCase{[]int{1, 5, 2, 3, 1}, 5},
		testCase{[]int{2, 1, 2, 0, 1}, 2},
	}
	ret := true
	for idx, testcase := range testCases {
		result := maxProfit(testcase.Prices)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %#v, %d, get %d\n", idx, testcase.Prices, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	ret := 0
	var currentBuy *int
	for idx := range prices {
		if currentBuy != nil {
			if (idx == len(prices)-1 || prices[idx+1] < prices[idx]) && prices[idx] > *currentBuy {
				ret += prices[idx] - *currentBuy
				// fmt.Printf("Sell at %d, earn %d\n", idx, prices[idx]-*currentBuy)
				currentBuy = nil
			}
		} else {
			if idx < len(prices)-1 && prices[idx+1] > prices[idx] {
				currentBuy = &prices[idx]
				// fmt.Printf("\nBuy at %d, price %d\n", idx, *currentBuy)
			}
		}
	}

	return ret
}

package solutions

import (
	"fmt"
	"go-leetcode-pratice/solutions/binarytreenode"
)

func AllPossibleFBT() bool {
	type testCase struct {
		N      int
		Expect string
	}

	testCases := []testCase{
		{3, "[0 0 <nil> <nil> 0 <nil> <nil>]"},
		{5, "[0 0 0 <nil> <nil> 0 <nil> <nil> 0 <nil> <nil>],[0 0 <nil> <nil> 0 0 <nil> <nil> 0 <nil> <nil>]"},
	}

	ret := true
	for idx, testcase := range testCases {
		var result binarytreenode.TreeNodes = allPossibleFBT(testcase.N)
		if result.ToString() != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: Input: %d\nExcept: %s, get %s\n", idx, testcase.N, testcase.Expect, result.ToString())
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

var cache = map[int][]*binarytreenode.TreeNode{}

func allPossibleFBT(n int) []*binarytreenode.TreeNode {
	ret := []*binarytreenode.TreeNode{}

	if cache[n] != nil {
		return cache[n]
	}

	if n < 1 {
		cache[n] = ret
		return ret
	}
	if n == 1 {
		ret = []*binarytreenode.TreeNode{
			{0, nil, nil},
		}
		cache[n] = ret
		return ret
	}

	for left := 1; left < n; left += 2 {
		leftChild := allPossibleFBT(left)
		rightChild := allPossibleFBT(n - 1 - left)
		for _, l := range leftChild {
			for _, r := range rightChild {
				ret = append(ret, &binarytreenode.TreeNode{
					Val:   0,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	cache[n] = ret
	return ret
}

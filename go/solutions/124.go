package solutions

import (
	"fmt"
	"math"
)

func MaxPathSum() bool {
	type testCase struct {
		Tree   *TreeNode
		Expect int
	}
	testCases := []testCase{
		testCase{sliceToTree([]interface{}{3, 1, 2}), 6},
		testCase{sliceToTree([]interface{}{-3}), -3},
		testCase{sliceToTree([]interface{}{-2, 6, nil, 0, -6}), 6},
		testCase{sliceToTree([]interface{}{2, -1}), 2},
		testCase{sliceToTree([]interface{}{1,
			nil, -7,
			nil, nil, -9, -8,
			nil, nil, nil, nil, nil, nil, 3, nil,
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, -2}), 3},
	}
	ret := true
	for idx, testcase := range testCases {
		result := maxPathSum2(testcase.Tree)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: Input: %s\nExcept: %d, get %d\n", idx, printTreeNode(testcase.Tree, 0), testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func maxPathSum(root *TreeNode) int {
	leftMax, leftCloseMax := maxSubTree(root.Left)
	rightMax, rightCloseMax := maxSubTree(root.Right)

	closeMax := root.Val
	if root.Left != nil && leftCloseMax > closeMax {
		closeMax = leftCloseMax
	}
	if root.Right != nil && rightCloseMax > closeMax {
		closeMax = rightCloseMax
	}

	max := root.Val
	if root.Left != nil {
		if leftMax > max {
			max = leftMax
		}
		if leftMax+root.Val > max {
			max = leftMax + root.Val
		}
	}
	if root.Right != nil {
		if rightMax > max {
			max = rightMax
		}
		if rightMax+root.Val > max {
			max = rightMax + root.Val
		}
		if root.Left != nil && rightMax+root.Val+leftMax > max {
			max = rightMax + root.Val + leftMax
		}
	}

	if max > closeMax {
		return max
	}
	return closeMax
}

func maxSubTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftMax, leftCloseMax := maxSubTree(root.Left)
	rightMax, rightCloseMax := maxSubTree(root.Right)

	closeMax := root.Val
	appendLeft := root.Val
	appendRight := root.Val
	if root.Left != nil {
		closeMax += leftMax
		appendLeft += leftMax
	}
	if root.Right != nil {
		closeMax += rightMax
		appendRight += rightMax
	}

	if root.Left != nil && leftCloseMax > closeMax {
		closeMax = leftCloseMax
	}
	if root.Right != nil && rightCloseMax > closeMax {
		closeMax = rightCloseMax
	}

	max := root.Val
	if appendLeft > max {
		max = appendLeft
	}

	if appendRight > max {
		max = appendRight
	}

	if max > closeMax {
		closeMax = max
	}
	// fmt.Printf("%#v, %d, %d\n", root, max, closeMax)
	return max, closeMax
}

func maxPathSum2(root *TreeNode) int {

	leftMax, leftCloseMax := maxSubTree2(root.Left)
	rightMax, rightCloseMax := maxSubTree2(root.Right)

	closeMax := root.Val + leftMax + rightMax
	if leftCloseMax > closeMax {
		closeMax = leftCloseMax
	}
	if rightCloseMax > closeMax {
		closeMax = rightCloseMax
	}

	max := root.Val
	if leftMax+root.Val > max {
		max = leftMax + root.Val
	}
	if rightMax+root.Val > max {
		max = rightMax + root.Val
	}

	if max > closeMax {
		return max
	}
	return closeMax
}

func maxSubTree2(root *TreeNode) (int, int) {
	if root == nil {
		return math.MinInt32, math.MinInt32
	}

	leftMax, leftCloseMax := maxSubTree2(root.Left)
	rightMax, rightCloseMax := maxSubTree2(root.Right)

	closeMax := root.Val + leftMax + rightMax
	if leftCloseMax > closeMax {
		closeMax = leftCloseMax
	}
	if rightCloseMax > closeMax {
		closeMax = rightCloseMax
	}

	max := root.Val
	if leftMax+root.Val > max {
		max = leftMax + root.Val
	}
	if rightMax+root.Val > max {
		max = rightMax + root.Val
	}

	if max > closeMax {
		closeMax = max
	}
	// fmt.Printf("%#v, %d, %d\n", root, max, closeMax)
	return max, closeMax
}

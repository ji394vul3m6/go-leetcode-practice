package solutions

import (
	"fmt"
)

func MinDiffInBST() bool {
	type testCase struct {
		Tree   *TreeNode
		Expect int
	}
	testCases := []testCase{
		testCase{sliceToTree([]interface{}{4, 2, 6, 1, 3, nil, nil}), 1},
		testCase{sliceToTree([]interface{}{90, 69, nil, 49, 89, nil, nil, nil, 52}), 1},
	}
	ret := true
	for idx, testcase := range testCases {
		result := minDiffInBST(testcase.Tree)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: Input: %s\nExcept: %d, get %d\n", idx, printTreeNode(testcase.Tree, 0), testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func minDiffInBST(root *TreeNode) int {
	slice := BSTToIntSlice(root)
	fmt.Printf("Get slice: [%v]\n", slice)
	if len(slice) <= 1 {
		// unexpected case
		return -1
	}
	min := slice[1] - slice[0]
	for idx := range slice {
		if idx+1 < len(slice) {
			diff := slice[idx+1] - slice[idx]
			if diff < min {
				min = diff
			}
		}
	}
	return min
}

func BSTToIntSlice(root *TreeNode) []int {
	ret := []int{}
	if root.Left != nil {
		ret = append(ret, BSTToIntSlice(root.Left)...)
	}
	ret = append(ret, root.Val)
	if root.Right != nil {
		ret = append(ret, BSTToIntSlice(root.Right)...)
	}
	return ret
}

/*
Given a Binary Search Tree (BST) with the root node `root`, return the minimum difference between the values of any two different nodes in the tree.

Example :

Input: root = [4,2,6,1,3,null,null]
Output: 1
Explanation:
Note that root is a TreeNode object, not an array.

The given tree [4,2,6,1,3,null,null] is represented by the following diagram:

          4
        /   \
      2      6
     /\
    1   3

while the minimum difference in this tree is 1, it occurs between node 1 and node 2, also between node 3 and node 2.

Note:

1.  The size of the BST will be between 2 and `100`.
2.  The BST is always valid, each node's value is an integer, and each node's value is different.
*/

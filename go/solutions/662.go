package solutions

import (
	"fmt"
)

func WidthOfBinaryTree() bool {
	type testCase struct {
		Tree   *TreeNode
		Expect int
	}
	testCases := []testCase{
		testCase{sliceToTree([]interface{}{1, 3, 2, 5, 3, nil, 9}), 4},
		testCase{sliceToTree([]interface{}{1, 3, nil, 5, 3}), 2},
		testCase{sliceToTree([]interface{}{1, 3, 2, 5}), 2},
		testCase{sliceToTree([]interface{}{1, 3, 2, 5, nil, nil, 9, 6, nil, nil, nil, nil, nil, nil, 7}), 8},
	}
	ret := true
	for idx, testcase := range testCases {
		result := widthOfBinaryTree(testcase.Tree)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: Input: %s\nExcept: %d, get %d\n", idx, printTreeNode(testcase.Tree, 0), testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

type MinMaxNode struct {
	Min int
	Max int
}

func widthOfBinaryTree(root *TreeNode) int {
	layerInfo := make(map[int]*MinMaxNode)
	currentLayer := 0

	travelNodes(root, currentLayer, 1, layerInfo)

	maxWidth := 0
	for _, layer := range layerInfo {
		if layer.Max-layer.Min > maxWidth {
			maxWidth = layer.Max - layer.Min
		}
	}

	return maxWidth + 1
}

func travelNodes(root *TreeNode, currentLayer int, pos int, layerInfo map[int]*MinMaxNode) {
	if _, ok := layerInfo[currentLayer]; ok {
		node := layerInfo[currentLayer]
		if pos < node.Min {
			node.Min = pos
		}
		if pos > node.Max {
			node.Max = pos
		}
	} else {
		layerInfo[currentLayer] = &MinMaxNode{pos, pos}
	}
	if root.Left != nil {
		travelNodes(root.Left, currentLayer+1, 2*pos-1, layerInfo)
	}
	if root.Right != nil {
		travelNodes(root.Right, currentLayer+1, 2*pos, layerInfo)
	}
	// fmt.Printf("Result at %d, pos=[%d, %d], node=[%v]\n", root.Val, currentLayer, pos, layerInfo[currentLayer])
}

/*
Given a binary tree, write a function to get the maximum width of the given tree. The width of a tree is the maximum width among all levels. The binary tree has the same structure as a full binary tree, but some nodes are null.

The width of one level is defined as the length between the end-nodes (the leftmost and right most non-null nodes in the level, where the `null` nodes between the end-nodes are also counted into the length calculation.

Example 1:

Input:

           1
         /   \
        3     2
       / \     \
      5   3     9

Output: 4
Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).

Example 2:

Input:

          1
         /
        3
       / \
      5   3

Output: 2
Explanation: The maximum width existing in the third level with the length 2 (5,3).

Example 3:

Input:

          1
         / \
        3   2
       /
      5

Output: 2
Explanation: The maximum width existing in the second level with the length 2 (3,2).

Example 4:

Input:

          1
         / \
        3   2
       /     \
      5       9
     /         \
    6           7
Output: 8
Explanation:The maximum width existing in the fourth level with the length 8 (6,null,null,null,null,null,null,7).

Note: Answer will in the range of 32-bit signed integer.
*/

package solutions

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToListNode(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	nodeList := []*ListNode{}
	for idx, n := range input {
		nodeList = append(nodeList, &ListNode{n, nil})
		if idx > 0 {
			nodeList[idx-1].Next = nodeList[idx]
		}
	}
	return nodeList[0]
}

func sliceToTree(input []interface{}) *TreeNode {
	nodes := make([]*TreeNode, len(input))
	for idx := range input {
		if w, ok := input[idx].(int); ok {
			nodes[idx] = &TreeNode{w, nil, nil}
		}
		if input[idx] != nil {
		}
	}
	for idx := range input {
		if nodes[idx] != nil {
			if 2*idx+1 < len(input) {
				nodes[idx].Left = nodes[2*idx+1]
			}
			if 2*idx+2 < len(input) {
				nodes[idx].Right = nodes[2*idx+2]
			}
		}
	}
	return nodes[0]
}

func testNodeSliceEqual(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	node1 := l1
	node2 := l2
	for node1 != nil && node2 != nil {
		if node1.Val != node2.Val {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return node1 == nil && node2 == nil
}

func printListNode(list *ListNode) string {
	if list == nil {
		return "nil"
	}
	node := list
	ret := []int{}
	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
	}
	return fmt.Sprint(ret)
}

func printTreeNode(tree *TreeNode, layer int) string {
	if tree == nil {
		return ""
	}
	left := printTreeNode(tree.Left, layer+1)
	right := printTreeNode(tree.Right, layer+1)

	spaces := ""
	for i := 0; i < layer; i++ {
		spaces = spaces + "   "
	}
	return fmt.Sprintf("%s\n%s%d\n%s", left, spaces, tree.Val, right)
}

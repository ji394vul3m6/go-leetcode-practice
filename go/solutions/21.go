package solutions

import (
	"fmt"
)

func MergeTwoLists() bool {
	type testCase struct {
		List1  *ListNode
		List2  *ListNode
		Expect *ListNode
	}
	testCases := []testCase{
		testCase{
			sliceToListNode([]int{1, 2, 4}),
			sliceToListNode([]int{1, 3, 4}),
			sliceToListNode([]int{1, 1, 2, 3, 4, 4}),
		},
	}
	ret := true
	for idx, testcase := range testCases {
		result := mergeTwoLists(testcase.List1, testcase.List2)
		if !testNodeSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail: %s, %s, %s, get %s\n",
				idx, printListNode(testcase.List1),
				printListNode(testcase.List2),
				printListNode(testcase.Expect),
				printListNode(result))
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := l1
	node2 := l2
	var ret *ListNode
	var current *ListNode
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	if node1.Val > node2.Val {
		ret = node2
		node2 = node2.Next
	} else {
		ret = node1
		node1 = node1.Next
	}

	current = ret
	for node1 != nil && node2 != nil {
		if node1.Val >= node2.Val {
			current.Next = node2
			current = node2
			node2 = node2.Next
		} else {
			current.Next = node1
			current = node1
			node1 = node1.Next
		}
	}

	for node1 != nil {
		current.Next = node1
		current = node1
		node1 = node1.Next
	}

	for node2 != nil {
		current.Next = node2
		current = node2
		node2 = node2.Next
	}

	return ret
}

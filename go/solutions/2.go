package solutions

import (
	"fmt"
)

func AddTwoNumbers() bool {
	type testCase struct {
		List1  *ListNode
		List2  *ListNode
		Expect *ListNode
	}
	testCases := []testCase{
		testCase{
			sliceToListNode([]int{2, 4, 3}),
			sliceToListNode([]int{5, 6, 4}),
			sliceToListNode([]int{7, 0, 8}),
		},
		testCase{
			sliceToListNode([]int{2, 4, 5, 1}),
			sliceToListNode([]int{5, 6, 6}),
			sliceToListNode([]int{7, 0, 2, 2}),
		},
		testCase{
			sliceToListNode([]int{5}),
			sliceToListNode([]int{5}),
			sliceToListNode([]int{0, 1}),
		},
		testCase{
			sliceToListNode([]int{1}),
			sliceToListNode([]int{9, 9, 9}),
			sliceToListNode([]int{0, 0, 0, 1}),
		},
	}
	ret := true
	for idx, testcase := range testCases {
		result := addTwoNumbers(testcase.List1, testcase.List2)
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := l1
	node2 := l2
	carry := false

	var ret []int
	for node1 != nil && node2 != nil {
		temp := node1.Val + node2.Val
		if carry {
			temp++
		}
		carry = temp >= 10
		ret = append(ret, temp%10)

		node1 = node1.Next
		node2 = node2.Next
	}

	remainNode := node1
	if remainNode == nil && node2 != nil {
		remainNode = node2
	}

	for remainNode != nil {
		temp := remainNode.Val
		if carry {
			temp++
		}
		carry = temp >= 10
		ret = append(ret, temp%10)
		remainNode = remainNode.Next
	}
	if carry {
		ret = append(ret, 1)
	}

	return sliceToListNode(ret)
}

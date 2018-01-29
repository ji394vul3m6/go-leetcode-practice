package solutions

import (
	"fmt"
)

func DeleteDuplicates() bool {
	type testCase struct {
		List   *ListNode
		Expect *ListNode
	}
	testCases := []testCase{
		testCase{
			sliceToListNode([]int{1, 1, 2}),
			sliceToListNode([]int{1, 2}),
		},
		testCase{
			sliceToListNode([]int{1, 1, 2, 3, 3}),
			sliceToListNode([]int{1, 2, 3}),
		},
	}
	ret := true
	for idx, testcase := range testCases {
		result := deleteDuplicates(testcase.List)
		if !testNodeSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail: %s, %s, get %s\n",
				idx, printListNode(testcase.List),
				printListNode(testcase.Expect),
				printListNode(result))
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	ret := []int{}
	ret = append(ret, head.Val)
	nowCheck := head.Next
	nowAddIdx := 0

	for nowCheck != nil {
		if nowCheck.Val != ret[nowAddIdx] {
			ret = append(ret, nowCheck.Val)
			nowAddIdx++
		}
		nowCheck = nowCheck.Next
	}

	return sliceToListNode(ret)
}

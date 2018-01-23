package solutions

import (
	"container/heap"
	"fmt"
)

func MergeKLists() bool {
	type testCase struct {
		List   []*ListNode
		Except *ListNode
	}
	testCases := []testCase{
		testCase{
			[]*ListNode{},
			nil,
		},
		testCase{
			[]*ListNode{sliceToListNode([]int{})},
			nil,
		},
		testCase{
			[]*ListNode{sliceToListNode([]int{}), sliceToListNode([]int{})},
			nil,
		},
		testCase{
			[]*ListNode{
				sliceToListNode([]int{1, 2, 5}),
				sliceToListNode([]int{1, 3, 4}),
			},
			sliceToListNode([]int{1, 1, 2, 3, 4, 5}),
		},
		testCase{
			[]*ListNode{
				sliceToListNode([]int{1, 5, 12}),
				sliceToListNode([]int{2, 6, 11}),
				sliceToListNode([]int{3, 7, 10}),
				sliceToListNode([]int{4, 8, 9}),
			},
			sliceToListNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
		},
	}
	ret := true
	for idx, testcase := range testCases {
		result := mergeKLists2(testcase.List)
		if !testNodeSliceEqual(result, testcase.Except) {
			ret = false
			showStr := ""
			for _, list := range testcase.List {
				showStr = showStr + " " + printListNode(list)
			}
			fmt.Printf("Case %d fail: %s, %s, get %s\n",
				idx, showStr,
				printListNode(testcase.Except),
				printListNode(result))
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func mergeKLists(lists []*ListNode) *ListNode {
	var ret *ListNode
	var current *ListNode

	if len(lists) == 0 {
		return nil
	}

	allNil := false
	for !allNil {
		minIdx := 0
		allNil = true
		for idx, list := range lists {
			allNil = allNil && (list == nil)
			if (list != nil && lists[minIdx] == nil) || (list != nil && list.Val < lists[minIdx].Val) {
				minIdx = idx
			}
		}

		if lists[minIdx] != nil {
			if ret == nil {
				ret = lists[minIdx]
				current = ret
			} else {
				current.Next = lists[minIdx]
				current = lists[minIdx]
			}

			lists[minIdx] = lists[minIdx].Next
		}
	}
	return ret
}

type ListHeap []*ListNode

func (h ListHeap) Len() int {
	return len(h)
}
func (h ListHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h ListHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *ListHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *ListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists2(lists []*ListNode) *ListNode {
	var ret *ListNode
	var current *ListNode
	h := &ListHeap{}
	for _, list := range lists {
		if list != nil {
			h.Push(list)
		}
	}
	heap.Init(h)
	for h.Len() > 0 {
		temp := heap.Pop(h).(*ListNode)

		if temp.Next != nil {
			h.Push(temp.Next)
			heap.Fix(h, len(*h)-1)
		}

		if ret == nil {
			ret = temp
			current = ret
		} else {
			current.Next = temp
			current = temp
			temp.Next = nil
		}
	}
	return ret
}

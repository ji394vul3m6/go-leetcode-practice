package solutions

import (
	"fmt"

	"go-leetcode-pratice/solutions/circularqueue"

	"go-leetcode-pratice/solutions/util"
)

func DesignCircularQueue() bool {
	type testCase struct {
		Size      int
		Operation []string
		Params    []int
		Expect    []interface{}
	}

	testCases := []testCase{
		testCase{
			3,
			[]string{"enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"},
			[]int{1, 2, 3, 4, 0, 0, 0, 4, 0},
			[]interface{}{true, true, true, false, 3, true, true, true, 4},
		},
		testCase{
			6,
			[]string{"enQueue", "Rear", "Rear", "deQueue", "enQueue", "Rear", "deQueue", "Front", "deQueue", "deQueue", "deQueue"},
			[]int{6, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0},
			[]interface{}{true, 6, 6, true, true, 5, true, -1, false, false, false},
		},
	}

	ret := true
	for idx, testcase := range testCases {
		result := runCircularQueue(testcase.Size, testcase.Operation, testcase.Params)
		if !util.TestSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail:%d\n%v\n%v\nExcept:\n%v\nGet:\n%v\n",
				idx, testcase.Size, testcase.Operation, testcase.Params, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func runCircularQueue(size int, operations []string, input []int) []interface{} {
	queue := circularqueue.Constructor(size)

	ret := []interface{}{}
	for idx, op := range operations {
		var val interface{}
		switch op {
		case "enQueue":
			val = queue.EnQueue(input[idx])
		case "deQueue":
			val = queue.DeQueue()
		case "Front":
			val = queue.Front()
		case "Rear":
			val = queue.Rear()
		case "isEmpty":
			val = queue.IsEmpty()
		case "isFull":
			val = queue.IsFull()
		default:
			fmt.Printf("Undefined op: %s\n", op)
			panic("")
		}
		ret = append(ret, val)
	}

	return ret
}

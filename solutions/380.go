package solutions

import (
	"fmt"

	"./randomizedcache"
	"./util"
)

func DesignRandomizedSet() bool {
	type testCase struct {
		Operation []string
		Params    []int
		Expect    []interface{}
	}

	testCases := []testCase{
		testCase{
			[]string{"insert", "remove", "insert", "get", "remove", "insert", "get"},
			[]int{1, 2, 2, 0, 1, 2, 0},
			[]interface{}{true, false, true, []interface{}{1, 2}, true, false, []interface{}{2}},
		},
	}

	ret := true
	for idx, testcase := range testCases {
		result := runRandomizedSet(testcase.Operation, testcase.Params)
		if !util.TestAnswerSlice(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail:%v\n%v\nExcept:\n%v\nGet:\n%v\n",
				idx, testcase.Operation, testcase.Params, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func runRandomizedSet(operations []string, input []int) []interface{} {
	set := randomizedcache.Constructor()

	ret := []interface{}{}
	for idx, op := range operations {
		var val interface{}
		switch op {
		case "insert":
			val = set.Insert(input[idx])
		case "remove":
			val = set.Remove(input[idx])
		case "get":
			val = set.GetRandom()
		default:
			fmt.Printf("Undefined op: %s\n", op)
			panic("")
		}
		ret = append(ret, val)
	}

	return ret
}

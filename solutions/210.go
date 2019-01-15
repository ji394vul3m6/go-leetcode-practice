package solutions

import (
	"fmt"
	"sort"

	"./treenode"
	"./util"
)

func FindOrder() bool {
	type testCase struct {
		Num            int
		PreRequirement [][]int
		Expect         []int
	}

	testCases := []testCase{
		testCase{2, [][]int{[]int{1, 0}}, []int{0, 1}},
		testCase{2, [][]int{[]int{1, 0}, []int{0, 1}}, []int{}},
		testCase{4, [][]int{[]int{1, 0}, []int{2, 0}, []int{3, 1}, []int{3, 2}}, []int{0, 1, 2, 3}},
	}

	ret := true
	for idx, testcase := range testCases {
		result := findOrder(testcase.Num, testcase.PreRequirement)
		if !util.TestIntSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d method1 fail: %#v, %d, %#v, get %#v\n", idx, testcase.Num, testcase.PreRequirement, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d method1 pass\n", idx)
		}
		result = findOrder2(testcase.Num, testcase.PreRequirement)
		if !util.TestIntSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d method2 fail: %#v, %d, %#v, get %#v\n", idx, testcase.Num, testcase.PreRequirement, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d method2 pass\n", idx)
		}
	}
	return ret
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	ret := []int{}
	pick := map[int]bool{}
	dependency := map[int][]int{}
	for _, prerequire := range prerequisites {
		if _, ok := dependency[prerequire[0]]; !ok {
			dependency[prerequire[0]] = []int{}
		}
		dependency[prerequire[0]] = append(dependency[prerequire[0]], prerequire[1])
	}

	for len(ret) < numCourses {
		pickOne := false
		for i := 0; i < numCourses; i++ {
			if pick[i] {
				continue
			}
			canPick := true
			for _, val := range dependency[i] {
				canPick = canPick && pick[val]
			}
			if canPick {
				pick[i] = true
				ret = append(ret, i)
				pickOne = true
			}
		}
		if !pickOne {
			return []int{}
		}
	}
	return ret
}

func findOrder2(numCourses int, prerequisites [][]int) []int {
	nodes := make(treenode.GraphNodes, numCourses)
	for idx := range nodes {
		nodes[idx].Value = idx
	}

	for _, prerequire := range prerequisites {
		target, front := prerequire[0], prerequire[1]
		nodes[front].Children = append(nodes[front].Children, target)
	}

	fill := 1
	cycle := false
	hasMeet := map[int]bool{}
	var DFSFill func(nodeIdx int)
	DFSFill = func(nodeIdx int) {
		if cycle {
			return
		}
		if nodes[nodeIdx].Order != 0 {
			return
		}
		hasMeet[nodeIdx] = true
		for idx := range nodes[nodeIdx].Children {
			val := nodes[nodeIdx].Children[len(nodes[nodeIdx].Children)-idx-1]
			if hasMeet[val] {
				cycle = true
				break
			}
			DFSFill(val)
		}
		nodes[nodeIdx].Order = fill
		fill++
		hasMeet[nodeIdx] = false
	}

	for idx := range nodes {
		DFSFill(idx)
	}
	if cycle {
		return []int{}
	}

	sort.Sort(nodes)
	ret := make([]int, len(nodes))
	for idx := range nodes {
		ret[idx] = nodes[idx].Value
	}

	return ret
}

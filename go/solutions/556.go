package solutions

import (
	"fmt"
	"math"
	"sort"
)

func NextGreaterElement() bool {
	type testCase struct {
		Input  int
		Expect int
	}
	testCases := []testCase{
		testCase{12, 21},
		testCase{21, -1},
		testCase{1356, 1365},
		testCase{1536, 1563},
		testCase{1273618273, 1273618327},
		testCase{1234507321, 1234510237},
		testCase{1234507123, 1234507132},
		testCase{1234507231, 1234507312},
		testCase{1999999999, -1},
	}
	ret := true
	for idx, testcase := range testCases {
		result := nextGreaterElement(testcase.Input)
		if result != testcase.Expect {
			ret = false
			fmt.Printf("Case %d fail: %d, %d, get %d\n", idx, testcase.Input, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func nextGreaterElement(n int) int {
	inputs := []int{}
	c := n
	for c != 0 {
		inputs = append(inputs, c%10)
		c /= 10
	}
	// fmt.Printf("Orig Inputs: [%v]\n", inputs)

	idx := 0
	for ; idx < len(inputs)-1; idx++ {
		if inputs[idx+1] < inputs[idx] {
			break
		}
	}
	if idx == len(inputs)-1 {
		return -1
	}

	changes := inputs[0 : idx+2]
	resort := []int{}
	candidates := []int{}
	// fmt.Printf("Changes: %#v\n", changes)
	for _, num := range changes {
		if num > inputs[idx+1] {
			candidates = append(candidates, num)
		} else {
			resort = append(resort, num)
		}
	}
	// fmt.Printf("Candidates: %v\nResort: %v\n", candidates, resort)
	sort.Ints(candidates)
	pick := candidates[0]
	resort = append(resort, candidates[1:]...)
	sort.Sort(sort.Reverse(sort.IntSlice(resort)))

	// fmt.Printf("Pick %d, Resort: [%v]\n", pick, resort)
	ret := append([]int{}, resort...)
	ret = append(ret, pick)
	ret = append(ret, inputs[idx+2:]...)
	// fmt.Printf("After Inputs: [%v]\n", ret)

	num := 0
	for i := len(ret) - 1; i >= 0; i-- {
		num *= 10
		num += ret[i]
	}

	if num > math.MaxInt32 {
		return -1
	}
	return num
}

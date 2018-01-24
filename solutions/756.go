package solutions

import (
	"fmt"
)

func PyramidTransition() bool {
	type testCase struct {
		Bottom  string
		Allowed []string
		Except  bool
	}
	testCases := []testCase{
		testCase{"XYZ", []string{"XYD", "YZE", "DEA", "FFF"}, true},
		testCase{"XXYX", []string{"XXX", "XXY", "XYX", "XYY", "YXZ"}, false},
		testCase{
			"CCC",
			[]string{"CBB", "ACB", "ABD", "CDB", "BDC", "CBC", "DBA", "DBB", "CAB", "BCB", "BCC", "BAA", "CCD", "BDD", "DDD", "CCA", "CAA", "CCC", "CCB"},
			true,
		},
		testCase{
			"BDBBAA",
			[]string{"ACB", "ACA", "AAA", "ACD", "BCD", "BAA", "BCB", "BCC", "BAD", "BAB", "BAC", "CAB", "CCD", "CAA", "CCA", "CCC", "CAD", "DAD", "DAA", "DAC", "DCD", "DCC", "DCA", "DDD", "BDD", "ABB", "ABC", "ABD", "BDB", "BBD", "BBC", "BBA", "ADD", "ADC", "ADA", "DDC", "DDB", "DDA", "CDA", "CDD", "CBA", "CDB", "CBD", "DBA", "DBC", "DBD", "BDA"},
			true,
		},
	}
	ret := true
	for idx, testcase := range testCases {
		result := pyramidTransition(testcase.Bottom, testcase.Allowed)
		if result != testcase.Except {
			ret = false
			fmt.Printf("Case %d fail: %s, %#v, %t, get %t\n", idx, testcase.Bottom, testcase.Allowed, testcase.Except, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func pyramidTransition(bottom string, allowed []string) bool {
	rules := make(map[string][]string)
	for _, pattern := range allowed {
		key := pattern[0:2]
		if _, ok := rules[key]; !ok {
			rules[key] = []string{}
		}
		rules[key] = append(rules[key], pattern[2:3])
	}

	working := [][]string{}
	for i := 0; i < len(bottom); i++ {
		working = append(working, []string{bottom[i : i+1]})
	}

	var nowWorking [][]string
	for len(working) > 1 {
		nowWorking = [][]string{}
		for i := 0; i < len(working)-1; i++ {
			part1 := working[i]
			part2 := working[i+1]

			temp := []string{}
			for _, p1 := range part1 {
				for _, p2 := range part2 {
					key := p1 + p2
					if vals, ok := rules[key]; ok {
						for _, val := range vals {
							if !containString(temp, val) {
								temp = append(temp, val)
							}
						}
					}
				}
			}
			if len(temp) == 0 {
				return false
			}
			nowWorking = append(nowWorking, temp)
		}
		working = nowWorking
	}
	return true
}

func containString(container []string, check string) bool {
	for _, v := range container {
		if v == check {
			return true
		}
	}
	return false
}

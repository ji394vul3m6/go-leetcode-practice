package solutions

import (
	"fmt"
	"strings"

	"go-leetcode-pratice/solutions/util"
)

func FindDuplicate() bool {
	type testCase struct {
		Files  []string
		Expect [][]string
	}
	testCases := []testCase{
		testCase{
			[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
			[][]string{
				[]string{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"},
				[]string{"root/a/1.txt", "root/c/3.txt"},
			},
		},
	}
	ret := true
	for idx, testcase := range testCases {
		result := findDuplicate(testcase.Files)
		if !util.TestDoubleStringSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail: %v\nExcept: %v\nGet: %v\n", idx, testcase.Files, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

func findDuplicate(paths []string) [][]string {
	contentFileMap := map[string][]string{}

	for _, p := range paths {
		params := strings.Split(p, " ")
		if len(params) <= 1 {
			continue
		}

		dir := params[0]
		files := params[1:]
		for _, name := range files {
			words := strings.Split(strings.Replace(name, ")", "", 1), "(")
			if len(words) != 2 {
				continue
			}

			filename := words[0]
			content := words[1]
			if _, ok := contentFileMap[content]; !ok {
				contentFileMap[content] = []string{}
			}
			contentFileMap[content] = append(contentFileMap[content], fmt.Sprintf("%s/%s", dir, filename))
		}
	}
	ret := [][]string{}
	for _, names := range contentFileMap {
		if len(names) <= 1 {
			continue
		}
		ret = append(ret, names)
	}
	return ret
}

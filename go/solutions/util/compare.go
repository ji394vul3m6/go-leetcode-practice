package util

import (
	"fmt"
	"sort"
	"strings"
)

func TestIntSliceEqual(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestStringSliceEqual(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestSliceEqual(a, b []interface{}) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestDoubleStringSliceEqual(a, b [][]string) bool {
	lines := []string{}
	for _, group := range a {
		sort.Strings(group)
		lines = append(lines, strings.Join(group, ","))
	}
	sort.Strings(lines)
	aStr := strings.Join(lines, ",")

	lines = []string{}
	for _, group := range b {
		sort.Strings(group)
		lines = append(lines, strings.Join(group, ","))
	}
	sort.Strings(lines)
	bStr := strings.Join(lines, ",")

	return aStr == bStr
}

func TestDoubleIntSliceEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	l := len(a)
	aStrs := make([]string, l)
	bStrs := make([]string, l)
	for idx := range a {
		sort.Ints(a[idx])
		aStr := ""
		for _, num := range a[idx] {
			aStr = fmt.Sprintf("%s,%d", aStr, num)
		}

		sort.Ints(b[idx])
		bStr := ""
		for _, num := range a[idx] {
			bStr = fmt.Sprintf("%s,%d", bStr, num)
		}

		aStrs = append(aStrs, aStr)
		bStrs = append(bStrs, bStr)
	}

	sort.Strings(aStrs)
	sort.Strings(bStrs)

	return strings.Join(aStrs, ".") == strings.Join(bStrs, ".")
}

func TestAnswerSlice(a, b []interface{}) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if candidate, ok := b[i].([]interface{}); ok {
			find := false
			for j := range candidate {
				if a[i] == candidate[j] {
					find = true
					break
				}
			}
			if !find {
				return false
			}
		} else {
			if a[i] != b[i] {
				return false
			}
		}
	}

	return true
}

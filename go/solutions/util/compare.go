package util

import (
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

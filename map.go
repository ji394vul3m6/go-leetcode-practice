package main

import "litttlebear/leetcode/solutions"

var (
	Questions map[int]interface{}
)

func init() {
	Questions = map[int]interface{}{
		1: solutions.TwoSum,
		2: solutions.AddTwoNumbers,
		3: solutions.LengthOfLongestSubstring,
	}
}

package main

import "litttlebear/leetcode/solutions"

var (
	Questions map[int]interface{}
)

func init() {
	Questions = map[int]interface{}{
		1:  solutions.TwoSum,
		2:  solutions.AddTwoNumbers,
		3:  solutions.LengthOfLongestSubstring,
		5:  solutions.LongestPalindrome,
		6:  solutions.ConvertZigZag,
		9:  solutions.IsPalindrome,
		20: solutions.IsValidParentheses,
		22: solutions.GenerateParenthesis,
		26: solutions.RemoveDuplicates,
	}
}

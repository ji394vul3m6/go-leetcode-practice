package main

import "litttlebear/leetcode/solutions"

var (
	Questions map[int]interface{}
)

func init() {
	Questions = map[int]interface{}{
		1:   solutions.TwoSum,
		2:   solutions.AddTwoNumbers,
		3:   solutions.LengthOfLongestSubstring,
		5:   solutions.LongestPalindrome,
		6:   solutions.ConvertZigZag,
		9:   solutions.IsPalindrome,
		17:  solutions.LetterCombinations,
		20:  solutions.IsValidParentheses,
		21:  solutions.MergeTwoLists,
		22:  solutions.GenerateParenthesis,
		23:  solutions.MergeKLists,
		26:  solutions.RemoveDuplicates,
		27:  solutions.RemoveSpecificElement,
		28:  solutions.ImplementStrStr,
		29:  solutions.SelfDivide,
		55:  solutions.CanJump,
		62:  solutions.UniquePaths,
		70:  solutions.ClimbStairs,
		83:  solutions.DeleteDuplicates,
		122: solutions.MaxProfit,
		124: solutions.MaxPathSum,
		146: solutions.RunLRUCacheTest,
		397: solutions.IntegerReplacement,
		755: solutions.PourWater,
		495: solutions.FindPoisonedDuration,
		658: solutions.FindClosestElements,
		756: solutions.PyramidTransition,
		783: solutions.MinDiffInBST,

		556: solutions.NextGreaterElement,
		227: solutions.Calculate,
		299: solutions.BullsandCows,
		662: solutions.WidthOfBinaryTree,
		766: solutions.IsToeplitzMatrix,
		692: solutions.TopKFrequent,
		598: solutions.RangeAdditionII,
	}
}

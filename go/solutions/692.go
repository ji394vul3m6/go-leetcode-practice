package solutions

import (
	"fmt"
	"strings"

	"go-leetcode-pratice/solutions/util"
)

func TopKFrequent() bool {
	type testCase struct {
		Words  []string
		K      int
		Expect []string
	}
	testCases := []testCase{
		testCase{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
		testCase{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4, []string{"the", "is", "sunny", "day"}},
		testCase{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 3, []string{"the", "is", "sunny"}},
		testCase{
			[]string{"glarko", "zlfiwwb", "nsfspyox", "pwqvwmlgri", "qggx", "qrkgmliewc", "zskaqzwo", "zskaqzwo", "ijy", "htpvnmozay", "jqrlad", "ccjel", "qrkgmliewc", "qkjzgws", "fqizrrnmif", "jqrlad", "nbuorw", "qrkgmliewc", "htpvnmozay", "nftk", "glarko", "hdemkfr", "axyak", "hdemkfr", "nsfspyox", "nsfspyox", "qrkgmliewc", "nftk", "nftk", "ccjel", "qrkgmliewc", "ocgjsu", "ijy", "glarko", "nbuorw", "nsfspyox", "qkjzgws", "qkjzgws", "fqizrrnmif", "pwqvwmlgri", "nftk", "qrkgmliewc", "jqrlad", "nftk", "zskaqzwo", "glarko", "nsfspyox", "zlfiwwb", "hwlvqgkdbo", "htpvnmozay", "nsfspyox", "zskaqzwo", "htpvnmozay", "zskaqzwo", "nbuorw", "qkjzgws", "zlfiwwb", "pwqvwmlgri", "zskaqzwo", "qengse", "glarko", "qkjzgws", "pwqvwmlgri", "fqizrrnmif", "nbuorw", "nftk", "ijy", "hdemkfr", "nftk", "qkjzgws", "jqrlad", "nftk", "ccjel", "qggx", "ijy", "qengse", "nftk", "htpvnmozay", "qengse", "eonrg", "qengse", "fqizrrnmif", "hwlvqgkdbo", "qengse", "qengse", "qggx", "qkjzgws", "qggx", "pwqvwmlgri", "htpvnmozay", "qrkgmliewc", "qengse", "fqizrrnmif", "qkjzgws", "qengse", "nftk", "htpvnmozay", "qggx", "zlfiwwb", "bwp", "ocgjsu", "qrkgmliewc", "ccjel", "hdemkfr", "nsfspyox", "hdemkfr", "qggx", "zlfiwwb", "nsfspyox", "ijy", "qkjzgws", "fqizrrnmif", "qkjzgws", "qrkgmliewc", "glarko", "hdemkfr", "pwqvwmlgri"},
			14,
			[]string{"nftk", "qkjzgws", "qrkgmliewc", "nsfspyox", "qengse", "htpvnmozay", "fqizrrnmif", "glarko", "hdemkfr", "pwqvwmlgri", "qggx", "zskaqzwo", "ijy", "zlfiwwb"},
		},
		testCase{
			[]string{"a", "b", "c", "a"},
			1,
			[]string{"a"},
		},
		testCase{
			[]string{"a", "b", "c", "c", "d"},
			4,
			[]string{"c", "a", "b", "d"},
		},
		testCase{
			[]string{"a", "b", "c", "b"},
			1,
			[]string{"b"},
		},
		testCase{
			[]string{"a", "a", "a", "b", "c", "b"},
			2,
			[]string{"a", "b"},
		},
		testCase{
			[]string{"a", "a", "a", "b", "c", "c"},
			2,
			[]string{"a", "c"},
		},
	}
	ret := true
	for idx, testcase := range testCases {
		result := topKFrequent(testcase.Words, testcase.K)
		if !util.TestStringSliceEqual(result, testcase.Expect) {
			ret = false
			fmt.Printf("Case %d fail: %v, %d\nExcept: %v\nGet   : %v\n", idx, testcase.Words, testcase.K, testcase.Expect, result)
		} else {
			fmt.Printf("Case %d pass\n", idx)
		}
	}
	return ret
}

type DoubleStringCountListNode struct {
	Count int
	Word  string
	Next  *DoubleStringCountListNode
	Last  *DoubleStringCountListNode
}

func topKFrequent(words []string, k int) []string {
	if len(words) == 0 {
		return []string{}
	}

	nodes := make(map[string]*DoubleStringCountListNode)

	start := &DoubleStringCountListNode{1, words[0], nil, nil}
	end := start
	nodes[words[0]] = start

	var insertNode *DoubleStringCountListNode
	var currentNode *DoubleStringCountListNode
	for _, word := range words[1:] {
		node, ok := nodes[word]
		if !ok {
			insertNode = &DoubleStringCountListNode{1, word, nil, nil}
			nodes[word] = insertNode
			currentNode = end
		} else if node == start {
			node.Count++
			continue
		} else {
			insertNode = node
			insertNode.Count++
			currentNode = insertNode.Last
			if insertNode.Last != nil {
				insertNode.Last.Next = insertNode.Next
			}
			if insertNode.Next != nil {
				insertNode.Next.Last = insertNode.Last
			}
			if insertNode == start && insertNode.Next != nil {
				start = insertNode.Next
			}
			if insertNode == end && insertNode.Last != nil {
				end = insertNode.Last
			}
			insertNode.Last = nil
			insertNode.Next = nil
		}

		for currentNode != nil &&
			(currentNode.Count < insertNode.Count ||
				(currentNode.Count == insertNode.Count && strings.Compare(currentNode.Word, insertNode.Word) > 0)) {
			currentNode = currentNode.Last
		}

		if currentNode == nil {
			// insert in first pos
			insertNode.Next = start
			start.Last = insertNode
			start = insertNode
		} else {
			if currentNode.Next != nil {
				currentNode.Next.Last = insertNode
				insertNode.Next = currentNode.Next
			}
			if currentNode == end {
				end = insertNode
			}
			currentNode.Next = insertNode
			insertNode.Last = currentNode
		}
	}

	ret := []string{}
	currentNode = start
	for currentNode != nil && len(ret) < k {
		ret = append(ret, currentNode.Word)
		currentNode = currentNode.Next
	}
	return ret
}

/*
Given a non-empty list of words, return the *k* most frequent elements.

Your answer should be sorted by frequency from highest to lowest. If two words have the same frequency, then the word with the lower alphabetical order comes first.

Example 1:

Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2

Output: ["i", "love"]

Explanation: "i" and "love" are the two most frequent words.

    Note that "i" comes before "love" due to a lower alphabetical order.

Example 2:

Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4

Output: ["the", "is", "sunny", "day"]

Explanation: "the", "is", "sunny" and "day" are the four most frequent words,

    with the number of occurrence being 4, 3, 2 and 1 respectively.

Note:

1\.  You may assume *k* is always valid, 1 ≤ *k* ≤ number of unique elements.

2\.  Input words contain only lowercase letters.

Follow up:

1\.  Try to solve it in *O*(*n* log *k*) time and *O*(*n*) extra space.
*/

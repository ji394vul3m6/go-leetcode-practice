package randomizedcache

import (
	"math/rand"
)

type RandomizedSet struct {
	Value map[int]bool
	Keys  []int
	Count int
}

// Constructor will Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		Value: map[int]bool{},
		Keys:  []int{},
	}
}

// Insert will Inserts a value to the set. Returns true if the set did not already contain the specified element.
func (set *RandomizedSet) Insert(val int) bool {
	if v, ok := set.Value[val]; ok {
		if v {
			return false
		}
	} else {
		set.Keys = append(set.Keys, val)
	}

	set.Value[val] = true
	set.Count++
	return true
}

// Remove will Removes a value from the set. Returns true if the set contained the specified element.
func (set *RandomizedSet) Remove(val int) bool {
	if !set.Value[val] {
		return false
	}
	set.Value[val] = false
	set.Count--

	if len(set.Keys) > 5*set.Count {
		keys := []int{}
		for key := range set.Value {
			keys = append(keys, key)
		}
		set.Keys = keys
	}

	return true
}

// GetRandom will Get a random element from the set.
func (set *RandomizedSet) GetRandom() int {
	if set.Count == 0 {
		return 0
	}
	idx := rand.Int() % len(set.Keys)

	for !set.Value[set.Keys[idx]] {
		idx = rand.Int() % len(set.Keys)
	}
	return set.Keys[idx]
}

package lrucache

import (
	"fmt"
)

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Last *DoubleListNode
}

type LRUCache struct {
	Pos      map[int]*DoubleListNode
	Value    map[int]int
	First    *DoubleListNode
	Last     *DoubleListNode
	Count    int
	MaxCount int
}

func Constructor(capacity int) LRUCache {
	ret := LRUCache{}
	ret.Pos = make(map[int]*DoubleListNode)
	ret.Value = make(map[int]int)
	ret.MaxCount = capacity
	return ret
}

func (this *LRUCache) putExistNodeToFirst(node *DoubleListNode) {
	if node.Last != nil {
		node.Last.Next = node.Next
	} else {
		// Last is nil, node is first node
		return
	}
	if node.Next != nil {
		node.Next.Last = node.Last
	} else {
		// Next is nil node is last node
		this.Last = node.Last
	}
	this.First.Last = node
	node.Next = this.First
	node.Last = nil
	this.First = node
}

func (this *LRUCache) putNewNodeToFirst(node *DoubleListNode) {
	if this.First != nil {
		this.First.Last = node
		node.Next = this.First
		this.First = node
	} else {
		this.First = node
		this.Last = node
	}
}

func (this *LRUCache) removeLastNode() {
	// remove last node
	current := this.Last
	if current.Last != nil {
		this.Last = current.Last
		current.Last.Next = nil
	} else {
		this.First = nil
		this.Last = nil
	}
	delete(this.Pos, current.Val)
}

func (this *LRUCache) Put(key int, value int) {
	if this.MaxCount == 0 {
		return
	}

	if node, ok := this.Pos[key]; ok {
		this.putExistNodeToFirst(node)
	} else {
		if this.Count >= this.MaxCount {
			this.removeLastNode()
		} else {
			this.Count++
		}
		new := &DoubleListNode{key, nil, nil}
		this.Pos[key] = new
		this.putNewNodeToFirst(new)
	}
	this.Value[key] = value
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Pos[key]; ok {
		this.putExistNodeToFirst(node)
		return this.Value[key]
	}
	return -1
}

func (this LRUCache) Debug() {
	// debugDoubleListNode(this.First, this.MaxCount)
	fmt.Printf("Key List: %s\n", printDoubleListNode(this.First, this.MaxCount))
	fmt.Printf("Values: %#v\n", this.Value)
	// fmt.Println("Pos:")
	// for k, v := range this.Pos {
	// 	fmt.Printf("\t%d: %p\n", k, v)
	// }
	// fmt.Printf("LRUCache: %#v\n", this)
	fmt.Println("========================")
}

func printDoubleListNode(list *DoubleListNode, max int) string {
	if list == nil {
		return "nil"
	}
	node := list
	ret := []int{}
	n := 1
	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
		n++
		if n > max && node != nil {
			fmt.Println("Error!")
			break
		}
	}
	return fmt.Sprint(ret)
}

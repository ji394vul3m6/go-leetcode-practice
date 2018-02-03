package solutions

import "fmt"

func RunLRUCacheTest() bool {
	type TestNode struct {
		Action string
		Key    int
		Value  int
	}
	type TestCase struct {
		Capacity int
		Nodes    []TestNode
	}

	tests := []TestCase{
		TestCase{
			2,
			[]TestNode{
				TestNode{"P", 1, 1},
				TestNode{"P", 2, 2},
				TestNode{"G", 1, 1},
				TestNode{"P", 3, 3}, // evicts key 2
				TestNode{"G", 2, -1},
				TestNode{"P", 4, 4}, // evicts key 1
				TestNode{"G", 1, -1},
				TestNode{"G", 3, 3},
				TestNode{"G", 4, 4},
			},
		},
		TestCase{
			2,
			[]TestNode{
				TestNode{"P", 2, 1},
				TestNode{"P", 2, 2},
				TestNode{"G", 2, 2},
				TestNode{"P", 1, 1}, // evicts key 2
				TestNode{"P", 4, 1}, // evicts key 1
				TestNode{"G", 2, -1},
			},
		},
		TestCase{
			3,
			[]TestNode{
				TestNode{"P", 1, 1},
				TestNode{"P", 2, 2},
				TestNode{"P", 3, 3},
				TestNode{"P", 4, 4},
				TestNode{"G", 1, -1},
				TestNode{"G", 2, 2},
				TestNode{"G", 2, 2},
				TestNode{"G", 3, 3},
			},
		},
		TestCase{
			5,
			[]TestNode{
				TestNode{"P", 1, 1},
				TestNode{"P", 2, 2},
				TestNode{"P", 3, 3},
				TestNode{"P", 4, 4},
				TestNode{"G", 1, 1},
				TestNode{"G", 2, 2},
				TestNode{"G", 2, 2},
				TestNode{"G", 3, 3},
			},
		},
		TestCase{
			10,
			[]TestNode{
				TestNode{"P", 10, 13},
				TestNode{"P", 3, 17},
				TestNode{"P", 6, 11},
				TestNode{"P", 10, 5},
				TestNode{"P", 9, 10},
				TestNode{"G", 13, -1},
				TestNode{"P", 2, 19},
				TestNode{"G", 2, 19},
				TestNode{"G", 3, 17},
				TestNode{"P", 5, 25},
				TestNode{"G", 8, -1},
				TestNode{"P", 9, 22},
				TestNode{"G", 9, 22},
				TestNode{"P", 5, 5},
				TestNode{"P", 1, 30},
				TestNode{"G", 11, -1},
				TestNode{"P", 9, 12},
				TestNode{"G", 7, -1},
				TestNode{"G", 5, 5},
				TestNode{"G", 8, -1},
				TestNode{"G", 9, 12},
				TestNode{"P", 4, 30},
				TestNode{"P", 9, 3},
				TestNode{"G", 9, 3},
				TestNode{"G", 10, 5},
				TestNode{"G", 10, 5},
				TestNode{"P", 6, 14},
				TestNode{"P", 3, 1},
				TestNode{"G", 3, 1},
				TestNode{"P", 10, 11},
				TestNode{"G", 8, -1},
				TestNode{"P", 2, 14},
				TestNode{"G", 1, 30},
				TestNode{"G", 5, 5},
				TestNode{"G", 4, 30},
				TestNode{"P", 11, 4},
				TestNode{"P", 12, 24},
				TestNode{"P", 5, 18},
				TestNode{"G", 13, -1},
				TestNode{"P", 7, 23},
				TestNode{"G", 8, -1},
				TestNode{"G", 12, 24},
				TestNode{"P", 3, 27},
				TestNode{"P", 2, 12},
				TestNode{"G", 5, 18},
				TestNode{"P", 2, 9},
				TestNode{"P", 13, 4},
				TestNode{"P", 8, 18},
				TestNode{"P", 1, 7},
				TestNode{"G", 6, -1},
				TestNode{"P", 9, 29},
				TestNode{"P", 8, 21},
				TestNode{"G", 5, 18},
				TestNode{"P", 6, 30},
				TestNode{"P", 1, 12},
				TestNode{"G", 10, -1},
				TestNode{"P", 4, 15},
				TestNode{"P", 7, 22},
				TestNode{"P", 11, 26},
				TestNode{"P", 8, 17},
				TestNode{"P", 9, 29},
				TestNode{"G", 5, 18},
				TestNode{"P", 3, 4},
				TestNode{"P", 11, 30},
				TestNode{"G", 12, -1},
				TestNode{"P", 4, 29},
				TestNode{"G", 3, 4},
				TestNode{"G", 9, 29},
				TestNode{"G", 6, 30},
				TestNode{"P", 3, 4},
				TestNode{"G", 1, 12},
				TestNode{"G", 10, -1},
				TestNode{"P", 3, 29},
				TestNode{"P", 10, 28},
				TestNode{"P", 1, 20},
				TestNode{"P", 11, 13},
				TestNode{"G", 3, 29},
				TestNode{"P", 3, 12},
				TestNode{"P", 3, 8},
				TestNode{"P", 10, 9},
				TestNode{"P", 3, 26},
				TestNode{"G", 8, 17},
				TestNode{"G", 7, 22},
				TestNode{"G", 5, 18},
				TestNode{"P", 13, 17},
				TestNode{"P", 2, 27},
				TestNode{"P", 11, 15},
				TestNode{"G", 12, -1},
				TestNode{"P", 9, 19},
				TestNode{"P", 2, 15},
				TestNode{"P", 3, 16},
				TestNode{"G", 1, 20},
				TestNode{"P", 12, 17},
				TestNode{"P", 9, 1},
				TestNode{"P", 6, 19},
				TestNode{"G", 4, -1},
				TestNode{"G", 5, 18},
				TestNode{"G", 5, 18},
				TestNode{"P", 8, 1},
				TestNode{"P", 11, 7},
				TestNode{"P", 5, 2},
				TestNode{"P", 9, 28},
				TestNode{"G", 1, 20},
				TestNode{"P", 2, 2},
				TestNode{"P", 7, 4},
				TestNode{"P", 4, 22},
				TestNode{"P", 7, 24},
				TestNode{"P", 9, 26},
				TestNode{"P", 13, 28},
				TestNode{"P", 11, 26},
			},
		},
	}

	for idx, test := range tests {
		capacity := test.Capacity
		cache := Constructor(capacity)
		fmt.Printf("\n\nTestcase %d start\n", idx)
		for _, node := range test.Nodes {
			if node.Action == "P" {
				cache.Put(node.Key, node.Value)

				fmt.Printf("Put %d %d\n", node.Key, node.Value)
				cache.Debug()
			} else if node.Action == "G" {
				ret := cache.Get(node.Key)
				fmt.Printf("Get %d %d\n", node.Key, node.Value)
				cache.Debug()

				if node.Value != ret {
					fmt.Printf("Get value of %d, expect %d, get %d\n", node.Key, node.Value, ret)
					fmt.Printf("Testcase %d fail\n", idx)
					return false
				}
			}
		}
	}
	return true
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

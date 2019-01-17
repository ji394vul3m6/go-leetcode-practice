package circularqueue

type MyCircularQueue struct {
	Space   []int
	HeadIdx int
	TailIdx int
	Length  int
}

// Constructor will initialize your data structure here. Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Space:   make([]int, k),
		HeadIdx: -1,
		TailIdx: -1,
		Length:  k,
	}
}

// EnQueue will insert an element into the circular queue. Return true if the operation is successful.
func (queue *MyCircularQueue) EnQueue(value int) bool {
	if queue.IsFull() {
		return false
	}
	if queue.IsEmpty() {
		queue.HeadIdx = 0
		queue.TailIdx = 0
	} else {
		queue.TailIdx = (queue.TailIdx + 1) % queue.Length
	}
	queue.Space[queue.TailIdx] = value
	return true
}

// DeQueue will delete an element from the circular queue. Return true if the operation is successful.
func (queue *MyCircularQueue) DeQueue() bool {
	if queue.IsEmpty() {
		return false
	}
	if queue.HeadIdx == queue.TailIdx {
		queue.HeadIdx = -1
		queue.TailIdx = -1
	} else {
		queue.HeadIdx = (queue.HeadIdx + 1) % queue.Length
	}
	return true
}

// Front will get the front item from the queue.
func (queue *MyCircularQueue) Front() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.Space[queue.HeadIdx]
}

// Rear will get the last item from the queue.
func (queue *MyCircularQueue) Rear() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.Space[queue.TailIdx]
}

// IsEmpty will checks whether the circular queue is empty or not.
func (queue *MyCircularQueue) IsEmpty() bool {
	return queue.TailIdx == -1 && queue.HeadIdx == -1
}

// IsFull will checks whether the circular queue is full or not.
func (queue *MyCircularQueue) IsFull() bool {
	return (queue.TailIdx+1)%queue.Length == queue.HeadIdx
}

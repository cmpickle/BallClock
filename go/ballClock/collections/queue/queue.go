package queue

import (
	"github.com/cmpickle/ballClock/go/ballClock/collections/node"
)

type (
	Queue struct {
		start, end *node.Node
		length     int
	}
)

// Create a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Take the next item off the front of the queue
func (this *Queue) Dequeue() *node.Node {
	if this.length == 0 {
		return nil
	}
	n := this.start
	if this.length == 1 {
		this.start = nil
		this.end = nil
	} else {
		this.start = this.start.Next
	}
	this.length--
	return n
}

// Put an item on the end of a queue
func (this *Queue) Enqueue(n *node.Node) {
	if this.length == 0 {
		this.start = n
		this.end = n
	} else {
		this.end.Next = n
		this.end = n
	}
	this.length++
}

// Put an item on the end of a queue
func (this *Queue) EnqueueN(nodes *node.Node, amount int) {
	if this.length == 0 {
		var end *node.Node
		for i := 0; i < amount; i++ {
			end = nodes.Next
		}
		this.start = nodes
		this.end = end
	} else {
		var end *node.Node
		for i := 0; i < amount; i++ {
			end = nodes.Next
		}
		this.end.Next = nodes
		this.end = end
	}
	this.length++
}

// Return the number of items in the queue
func (this *Queue) Len() int {
	return this.length
}

// Return the first item in the queue without removing it
func (this *Queue) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.start.Value
}

func (this *Queue) ElementAt(pos int) node.Node {
	curr := this.start

	for i := 0; i < pos; i++ {
		curr = curr.Next
	}

	return *curr
}

func (this *Queue) ToArray() []int {
	a := make([]int, this.Len())
	for i := 0; i < this.Len(); i++ {
		a[i] = this.ElementAt(i).Value.(int)
	}
	return a
}

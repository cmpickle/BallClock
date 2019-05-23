package stack

import (
	"github.com/cmpickle/ballClock/go/ballClock/collections/node"
)

type (
	Stack struct {
		top    *node.Node
		length int
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.Value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() *node.Node {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.Next
	this.length--
	return n
}

// Pop the top item of the stack and return it
func (this *Stack) PopN(amount int) *node.Node {
	if this.length < amount {
		return nil
	}
	var result *node.Node
	result = this.top

	for i := 0; i < amount; i++ {
		this.top = this.top.Next
	}
	this.length -= amount
	return result
}

// Push a value onto the top of the stack
func (this *Stack) Push(n *node.Node) {
	n.Next = this.top
	this.top = n
	this.length++
}

func (this *Stack) ElementAt(pos int) node.Node {
	curr := this.top

	for i := 0; i < pos; i++ {
		curr = curr.Next
	}

	return *curr
}

func (this *Stack) ToArray() []int {
	a := make([]int, this.Len())
	for i := 0; i < this.Len(); i++ {
		a[this.Len()-1-i] = this.ElementAt(i).Value.(int)
	}
	return a
}

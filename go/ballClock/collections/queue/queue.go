package queue

type (
	Queue struct {
		start, end *node
		length     int
	}
	node struct {
		Value interface{}
		next  *node
	}
)

// Create a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Take the next item off the front of the queue
func (this *Queue) Dequeue() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.start
	if this.length == 1 {
		this.start = nil
		this.end = nil
	} else {
		this.start = this.start.next
	}
	this.length--
	return n.Value
}

// Put an item on the end of a queue
func (this *Queue) Enqueue(Value interface{}) {
	n := &node{Value, nil}
	if this.length == 0 {
		this.start = n
		this.end = n
	} else {
		this.end.next = n
		this.end = n
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

func (this *Queue) ElementAt(pos int) node {
	curr := this.start

	for i := 0; i < pos; i++ {
		curr = curr.next
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

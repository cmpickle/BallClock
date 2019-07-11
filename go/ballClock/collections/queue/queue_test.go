package queue

import (
	"reflect"
	"testing"

	"github.com/cmpickle/ballClock/go/ballClock/collections/node"
)

func Test(t *testing.T) {
	q := New()

	if q.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	q.Enqueue(&node.Node{1, nil})

	if q.Len() != 1 {
		t.Errorf("Length should be 1")
	}

	if q.Peek().(int) != 1 {
		t.Errorf("Enqueued value should be 1")
	}

	v := q.Dequeue()

	if v.Value.(int) != 1 {
		t.Errorf("Dequeued value should be 1")
	}

	if q.Peek() != nil || q.Dequeue() != nil {
		t.Errorf("Empty queue should have no values")
	}

	q.Enqueue(&node.Node{1, nil})
	q.Enqueue(&node.Node{2, nil})

	if q.Len() != 2 {
		t.Errorf("Length should be 2")
	}

	if q.Peek().(int) != 1 {
		t.Errorf("First value should be 1")
	}

	q.Dequeue()

	if q.Peek().(int) != 2 {
		t.Errorf("Next value should be 2")
	}
}

func TestElementAt(t *testing.T) {
	s := New()

	s.Enqueue(&node.Node{1, nil})
	s.Enqueue(&node.Node{2, nil})

	if s.ElementAt(0).Value.(int) != 1 {
		t.Errorf("Element 0: expected 1 actual %v", s.ElementAt(0).Value.(int))
	}

	if s.ElementAt(1).Value.(int) != 2 {
		t.Errorf("Element 1: expected 2 actual %v", s.ElementAt(1).Value.(int))
	}
}

func TestToArray(t *testing.T) {
	s := New()

	s.Enqueue(&node.Node{1, nil})
	s.Enqueue(&node.Node{2, nil})

	actual := s.ToArray()
	expected := []int{1, 2}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v actual %v", expected, actual)
	}
}

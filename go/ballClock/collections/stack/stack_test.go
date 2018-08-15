package stack

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	s := New()

	if s.Len() != 0 {
		t.Errorf("Length of an empty stack should be 0")
	}

	s.Push(1)

	if s.Len() != 1 {
		t.Errorf("Length should be 0")
	}

	if s.Peek().(int) != 1 {
		t.Errorf("Top item on the stack should be 1")
	}

	if s.Pop().(int) != 1 {
		t.Errorf("Top item should have been 1")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	s.Push(1)
	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("Length should be 2")
	}

	if s.Peek().(int) != 2 {
		t.Errorf("Top of the stack should be 2")
	}
}

func TestPeekNil(t *testing.T) {
	s := New()

	actual := s.Peek()

	if actual != nil {
		t.Errorf("Expected return of nil from empty Peek, actual %v", actual)
	}
}

func TestPopNil(t *testing.T) {
	s := New()

	actual := s.Pop()

	if actual != nil {
		t.Errorf("Expected return of nil from empty Pop, actual %v", actual)
	}
}

func TestElementAt(t *testing.T) {
	s := New()

	s.Push(1)
	s.Push(2)

	if s.ElementAt(0).value.(int) != 2 {
		t.Errorf("Element 0: expected 2 actual %v", s.ElementAt(0).value.(int))
	}

	if s.ElementAt(1).value.(int) != 1 {
		t.Errorf("Element 1: expected 1 actual %v", s.ElementAt(1).value.(int))
	}
}

func TestToArray(t *testing.T) {
	s := New()

	s.Push(1)
	s.Push(2)

	actual := s.ToArray()
	expected := []int{1, 2}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v actual %v", expected, actual)
	}
}

package BallClock

import (
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

type BallClock struct {
	Count   int
	Min     *stack.Stack
	FiveMin *stack.Stack
	Hour    *stack.Stack
	Main    *queue.Queue
}

func New(count int) *BallClock {
	return &BallClock{count, stack.New(), stack.New(), stack.New(), queue.New()}
}

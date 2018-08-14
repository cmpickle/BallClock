package BallClock

import (
	"github.com/cmpickle/ballClock/go/ballClock/collections"
	"github.com/golang-collections/collections/stack"
)

type BallClock struct {
	Count   int
	Min     *stack.Stack
	FiveMin *stack.Stack
	Hour    *stack.Stack
	Main    *collections.Queue
}

func New(count int) *BallClock {
	return &BallClock{count, stack.New(), stack.New(), stack.New(), collections.New()}
}

func (ballClock *BallClock) Tick() {
	ball := ballClock.Main.Dequeue()
	if ball == nil {
		return
	}

	if ballClock.Min.Len() < 4 {
		ballClock.Min.Push(ball)
	} else {
		ballClock.returnBalls(ballClock.Min)

		if ballClock.FiveMin.Len() < 11 {
			ballClock.FiveMin.Push(ball)
		} else {
			ballClock.returnBalls(ballClock.FiveMin)

			if ballClock.Hour.Len() < 11 {
				ballClock.Hour.Push(ball)
			} else {
				ballClock.returnBalls(ballClock.Hour)

				ballClock.Main.Enqueue(ball)
			}
		}
	}
}

func (ballClock *BallClock) returnBalls(stack *stack.Stack) {
	for stack.Len() > 0 {
		ball := stack.Pop()
		if ball != nil {
			ballClock.Main.Enqueue(ball)
		}
	}
}

func (ballClock *BallClock) isStartingOrder() bool {
	arr := ballClock.Main.ToArray()
	for i := 0; i < ballClock.Main.Len(); i++ {
		if arr[i] != i+1 {
			return false
		}
	}
	return true
}

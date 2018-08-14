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

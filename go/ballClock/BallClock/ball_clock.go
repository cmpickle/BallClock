package ballclock

import (
	"encoding/json"
	"fmt"

	"github.com/cmpickle/ballClock/go/ballClock/collections/queue"
	"github.com/cmpickle/ballClock/go/ballClock/collections/stack"
)

type BallClock struct {
	Count   int
	Min     *stack.Stack `json:"Min"`
	FiveMin *stack.Stack `json:"FiveMin"`
	Hour    *stack.Stack `json:"Hour"`
	Main    *queue.Queue `json:"Main"`
}

type BallClockJson struct {
	Min     []int `json:"Min"`
	FiveMin []int `json:"FiveMin"`
	Hour    []int `json:"Hour"`
	Main    []int `json:"Main"`
}

func New(count int) *BallClock {
	ballClock := &BallClock{count, stack.New(), stack.New(), stack.New(), queue.New()}
	ballClock.Main.Init(count)
	// for i := 1; i <= count; i++ {
	// 	ballClock.Main.Enqueue(&node.Node{i, nil})
	// }
	return ballClock
}

func (ballClock *BallClock) Ticks(minutes int) {
	for i := 0; i < minutes; i++ {
		ballClock.Tick()
	}
}

func (ballClock *BallClock) Tick() {
	ball := ballClock.Main.Dequeue()

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
	// for stack.Len() > 0 {
	var length = stack.Len()
	balls := stack.PopN(length)
	if balls != nil {
		ballClock.Main.EnqueueN(balls, length)
	}
	// }
}

func (ballClock *BallClock) IsStartingOrder() bool {
	for i := 0; i < ballClock.Main.Len()-1; i++ {
		if ballClock.Main.ElementAt(i).Value != i+1 {
			return false
		}
	}
	return true
}

func (ballClock *BallClock) ToJson() []byte {
	bc := &BallClockJson{Main: ballClock.Main.ToArray(), FiveMin: ballClock.FiveMin.ToArray(), Hour: ballClock.Hour.ToArray(), Min: ballClock.Min.ToArray()}
	result, err := json.Marshal(bc)
	if err != nil {
		fmt.Println("Error parsing Ball Clock to Json")
	}
	return result
}

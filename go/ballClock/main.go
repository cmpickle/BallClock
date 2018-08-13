package main

import (
	"fmt"

	"github.com/cmpickle/ballClock/go/ballClock/BallClock"
)

func main() {
	fmt.Println("Hello World")
	var ballClock *BallClock.BallClock = BallClock.New(30)
	fmt.Println(ballClock.Count)
}

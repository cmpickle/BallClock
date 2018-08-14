package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cmpickle/ballClock/go/ballClock/BallClock"
)

func main() {
	var ballClock *BallClock.BallClock

	switch os.Args[1] {
	case "--repetitionTime":
		balls, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Errorf("repetition time requires an integer parameter between 27 and 127.")
		}
		ballClock = BallClock.New(balls)
		days := 0
		for {
			for i := 0; i < 1440; i++ {
				ballClock.Tick()
			}
			days++
			if ballClock.IsStartingOrder() {
				break
			}
		}

		fmt.Printf("%v balls cycle after %v days.", balls, days)
	}

	// fmt.Println(ballClock.Count)
	// for i := 0; i < 1440*15; i++ {
	// 	ballClock.Tick()
	// }
	// fmt.Println(ballClock.IsStartingOrder())
}

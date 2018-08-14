package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cmpickle/ballClock/go/ballClock/BallClock"
)

func main() {
	var ballClock *BallClock.BallClock

	if len(os.Args) <= 1 {
		fmt.Printf("Please enter a valid command line option:\n\t--repetitionTime\n\t--timedOutput")
		return
	}
	switch os.Args[1] {
	case "--repetitionTime":
		balls, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(fmt.Errorf("repetition time requires an integer parameter between 27 and 127."))
			return
		}
		if balls < 27 || balls > 127 {
			fmt.Println(fmt.Errorf("repetition time requires an integer parameter between 27 and 127."))
			return
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
		break
	default:
		fmt.Printf("Please enter a valid command line option:\n\t--repetitionTime\n\t--timedOutput")
	}

	// fmt.Println(ballClock.Count)
	// for i := 0; i < 1440*15; i++ {
	// 	ballClock.Tick()
	// }
	// fmt.Println(ballClock.IsStartingOrder())
}

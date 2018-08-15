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
			ballClock.Ticks(1440)
			days++
			if ballClock.IsStartingOrder() {
				break
			}
		}

		fmt.Printf("%v balls cycle after %v days.", balls, days)
		break
	case "--timedOutput":
		balls, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(fmt.Errorf("repetition time requires an integer parameter balls between 27 and 127."))
			return
		}
		if balls < 27 || balls > 127 {
			fmt.Println(fmt.Errorf("repetition time requires an integer parameter balls between 27 and 127."))
			return
		}
		minutes, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(fmt.Errorf("repetition time requires an integer parameter minutes greater than 0."))
			return
		}
		if balls < 27 || balls > 127 {
			fmt.Println(fmt.Errorf("repetition time requires an integer parameter minutes greater than 0."))
			return
		}
		ballClock = BallClock.New(balls)
		ballClock.Ticks(minutes)
		fmt.Println(string(ballClock.ToJson()))
	default:
		fmt.Printf("Please enter a valid command line option:\n\t--repetitionTime\n\t--timedOutput")
	}
}

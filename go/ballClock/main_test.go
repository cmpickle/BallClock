package main

import (
	"testing"
	"github.com/cmpickle/ballClock/go/ballClock/BallClock"
)

func BenchmarkTicks10(b *testing.B) {
	var ballClock *BallClock.BallClock
	ballClock = BallClock.New(123)
	days := 0
	for {
		ballClock.Ticks(1440)
		days++
		if ballClock.IsStartingOrder() {
			break
		}
	}
}

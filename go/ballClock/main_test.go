package main

import (
	"testing"

	"cameronpickle.com/ballClock/Clock"
)

func BenchmarkTicks10(b *testing.B) {
	ballClock := Clock.New(123)
	days := 0
	for {
		ballClock.Ticks(1440)
		days++
		if ballClock.IsStartingOrder() {
			break
		}
	}
}

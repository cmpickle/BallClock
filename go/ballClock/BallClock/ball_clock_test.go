package BallClock

import (
	"testing"
)

func TestNew(t *testing.T) {
	bc := New(30)

	if bc.Count != 30 {
		t.Errorf("Expected count to be 30, actual %v", bc.Count)
	}
}

func TestTick(t *testing.T) {
	bc := New(30)

	bc.Tick()

	actual := bc.Min.Pop()

	if actual.Value != 1 {
		t.Errorf("Expected 1 to be in minute stack after Tick, actual %v", actual)
	}
}

func TestTicks(t *testing.T) {
	bc := New(30)

	bc.Ticks(65)

	actualFiveMin := bc.FiveMin.Pop()
	actualHour := bc.Hour.Pop()

	if actualFiveMin.Value != 11 {
		t.Errorf("Expected 5 to be in five minute stack after 65 Ticks, actual %v", actualFiveMin)
	}

	if actualHour.Value != 6 {
		t.Errorf("Expected 11 to be in hour stack after 65 Ticks, actual %v", actualHour)
	}
}

func TestIsStartingOrder27Balls23Days(t *testing.T) {
	bc := New(27)

	bc.Ticks(23 * 1440)

	if !bc.IsStartingOrder() {
		t.Errorf("Expected 27 balls to be in starting order after 23 days")
	}
}

func TestIsStartingOrder30Balls15Days(t *testing.T) {
	bc := New(30)

	bc.Ticks(15 * 1440)

	if !bc.IsStartingOrder() {
		t.Errorf("Expected 30 balls to be in starting order after 15 days")
	}
}

func TestIsStartingOrder45Balls378Days(t *testing.T) {
	bc := New(45)

	bc.Ticks(378 * 1440)

	if !bc.IsStartingOrder() {
		t.Errorf("Expected 45 balls to be in starting order after 378 days")
	}
}

func TestIsStartingOrderFalse(t *testing.T) {
	bc := New(30)

	bc.Tick()

	if bc.IsStartingOrder() {
		t.Errorf("Expected 30 balls to be in not be in starting order after 1 minute")
	}
}

func TestToJson(t *testing.T) {
	bc := New(30)

	bc.Ticks(75)

	actual := string(bc.ToJson())
	expected := `{"Min":[],"FiveMin":[11,16,14],"Hour":[6],"Main":[4,3,2,1,21,17,13,9,30,25,20,15,10,5,19,18,8,7,24,23,22,12,29,28,27,26]}`

	if actual != expected {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

var result int

func BenchmarkIsStartingOrder(b *testing.B) {
	bc := New(27)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bc.IsStartingOrder()
	}
}

func BenchmarkReturnBalls(b *testing.B) {
	bc := New(30)
	bc.Ticks(4)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bc.returnBalls(bc.Min)
	}
}

func BenchmarkTick(b *testing.B) {
	bc := New(30)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bc.Tick()
	}
}

func BenchmarkTicks1440(b *testing.B) {
	bc := New(30)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bc.Ticks(1440)
	}
}

package calculator

import (
	"testing"
	"time"
)

func TestCalculator(t *testing.T) {
	input := make(chan int)
	output := make(chan int)
	calculator := Calculator{
		Input:  input,
		Output: output,
	}
	calculator.Start()
	for i := 0; i < 100; i++ {
		input <- i
		got := <-output
		expected := i * i
		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	}
	close(input)
	time.Sleep(time.Millisecond) // give some time to get closed
	select {
	case _, ok := <-output:
		if ok {
			t.Fatal("output channel must be closed after input channel was closed")
		}
	default:
		t.Fatal("output channel must be closed after input channel was closed")
	}
}

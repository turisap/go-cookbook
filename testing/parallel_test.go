package kirmath

import (
	"fmt"
	"testing"
)

func TestAddOneDigit(t *testing.T) {
	t.Parallel()
	result := Add(1, 2)
	if result != 3 {
		t.Error("Adding 1 and 2 doesn't produce 3")
	}
}
func TestAddTwoDigits(t *testing.T) {
	t.Parallel()
	result := Add(12, 30)
	if result != 42 {
		t.Error("Adding 12 and 30 doesn't produce 42")
	}
}
func TestAddThreeDigits(t *testing.T) {
	t.Parallel()
	result := Add(100, -1)
	if result != 99 {
		t.Error("Adding 100 and -1 doesn't produce 99")
	}
}

func FuzzGuess(f *testing.F) {
	f.Fuzz(func(t *testing.T, input int) {
		fmt.Println("INPUT", input)
		Guess(input)
	})
}

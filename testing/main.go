package kirmath

import "time"

func Add(x, y int) int {
	time.Sleep(500 * time.Millisecond)
	return x + y
}

func Guess(n int) {
	if n == 25 {
		panic("blackjack!")
	}

}

package kirmath

func Add(x, y int) int {
	return x + y
}

func Guess(n int) {
	if n == 25 {
		panic("blackjack!")
	}

}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

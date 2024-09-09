package pkg

// Fibonacci is a function that returns the n-th number of the Fibonacci sequence.
// 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 ...
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

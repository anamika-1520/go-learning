package main

type std struct {
	Name    string
	Age     int
	Class   string
	Roll_no int
	Marks   int
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

package main

import "fmt"

func fact(n int) int {

	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)
	fmt.Println(fact(n))
}

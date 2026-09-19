package main

import "fmt"

func main() {

	marks := [5]int{80, 75, 90, 85, 70}

	for i := 0; i < len(marks); i++ {
		fmt.Print(marks[i], " ")
	}
	fmt.Print("\n")
	fmt.Println(marks)
}

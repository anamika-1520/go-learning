package main

import "fmt"

func main() {

	marks := 85

	if marks >= 90 {
		fmt.Println("A+")
	} else if marks >= 80 {
		fmt.Println("A")
	} else if marks >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("Fail")
	}
}

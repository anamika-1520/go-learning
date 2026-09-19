package main

import "fmt"

func main() {
	var day int
	fmt.Print("Enter a day number (1-7): ")
	fmt.Scanf("%d", &day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	var marks int
	fmt.Print("Enter your marks: ")
	fmt.Scanf("%d", &marks)
	switch {
	case marks >= 90:
		fmt.Println("A+")
	case marks >= 80:
		fmt.Println("A")
	case marks >= 70:
		fmt.Println("B")
	default:
		fmt.Println("Fail")
	}

}

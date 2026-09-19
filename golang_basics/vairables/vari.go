package main

import "fmt" // # its a global variable
var name string = "saloni"

func main() {
	name := "Anamika" // # its a local variable
	age := 23
	salary := 40000.5
	isWorking := true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(salary)
	fmt.Println(isWorking)
}

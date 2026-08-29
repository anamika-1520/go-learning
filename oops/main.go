package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("fact of %v is %v", a, factorial(a))

	var s student
	s.Name = "John"
	s.Age = 20
	s.Class = "12th"
	s.Roll_no = 101
	s.Marks = 85

	fmt.Printf("\nStudent Details:\n")
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Class: %s\n", s.Class)
	fmt.Printf("Roll No: %d\n", s.Roll_no)
	fmt.Printf("Marks: %d\n", s.Marks)
	s.display()
	s.updateAge(25)
	s.updateClass("bsc")
	s.updateName("Anamika")
	s.updateMarks(123)
	s.display()
}

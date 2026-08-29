package main

import "fmt"

type student struct {
	Name    string
	Age     int
	Class   string
	Roll_no int
	Marks   int
}

func (s student) display() {
	fmt.Println("\nStudent Details:")
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Class: %s\n", s.Class)
	fmt.Printf("Roll No: %d\n", s.Roll_no)
	fmt.Printf("Marks: %d\n", s.Marks)

}
func (s student) updateMarks(newMarks int) { // marks globally update nhi hogi
	s.Marks = newMarks
	fmt.Printf("\nUpdated Marks: %d\n", s.Marks)
}
func (s student) updateAge(newAge int) { // age globally update nhi hogi
	s.Age = newAge
	fmt.Printf("\nUpdated Age: %d\n", s.Age)
}
func (s *student) updateName(newName string) {
	s.Name = newName
	fmt.Printf("\nUpdated Name: %s\n", s.Name)
}
func (s *student) updateClass(newClass string) {
	s.Class = newClass
	fmt.Printf("\nUpdated Class: %s\n", s.Class)
}

package main

import "fmt"

type Student struct {
	name  string
	marks int
}

func main() {
	m := make(map[string]int)
	for i := 0; i < 5; i++ {
		var name string
		var marks int
		fmt.Print("Enter name: ")
		fmt.Scanf("%s", &name)
		fmt.Print("Enter marks: ")
		fmt.Scanf("%d", &marks)
		m[name] = marks
	}
	fmt.Println(m)
	s := make(map[string]Student)
	for i := 0; i < 5; i++ {
		var name string
		var marks int
		fmt.Print("Enter name: ")
		fmt.Scanf("%s", &name)
		fmt.Print("Enter marks: ")
		fmt.Scanf("%d", &marks)
		s[name] = Student{name, marks} // # creating a struct object and storing it in the map
	}
	fmt.Println(s)
}

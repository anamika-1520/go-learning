package main

import "fmt"

type Runner interface {
	Run()
}

type Dog struct{}

func (d Dog) Run() {
	fmt.Println("Dog is running")
}

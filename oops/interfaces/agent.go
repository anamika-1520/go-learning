package main

import "fmt"

type agent interface {
	// method signatures
	// methods that an agent should implement
	// for example:
	// getName() string
	// getID() int
	Run()           // method signature for running
	Stop()          // method signature for stopping
	Status() string // method signature for getting status

}

type Chatagent struct {
	// fields for Chatagent
}
type Voiceagent struct {
}
type Emailagent struct {
}

func NewChatagent() *Chatagent {
	fmt.Println("Chatagent is running")
	return &Chatagent{}
}

func (c *Chatagent) Run() {
	fmt.Println("chatagnt is running !!!!!!")
}
func (c *Chatagent) Stop() {
	fmt.Println("chatagnt is stopped !!!!!!")
}
func (c *Chatagent) Status() string {
	return "Chatagent is active"
}

func (v Voiceagent) Run() {
	fmt.Println("Voiceagent is running !!!!!!")
}
func (v Voiceagent) Stop() {
	fmt.Println("Voiceagent is stopped !!!!!!")
}
func (v Voiceagent) Status() string {
	return "Voiceagent is active"
}

func (e Emailagent) Run() {
	fmt.Println("Emailagent is running !!!!!!")

}
func (e Emailagent) Stop() {
	fmt.Println("email agent is stopped. !!!!")

}
func (e Emailagent) Status() string {
	return "Emailagent is active"
}

package main

import "fmt"

func main() {
	// use interfaces.Runner
	var d Dog
	d.Run()
	var chat Chatagent
	chat.Run()
	chat.Stop()
	fmt.Println(chat.Status())
	var voice Voiceagent
	voice.Run()
	voice.Stop()
	fmt.Println(voice.Status())
	var email Emailagent
	email.Run()
	email.Stop()
	fmt.Println(email.Status())
}

package main

import "fmt"

func main() {
	greetings := []string{
		"Hello",
		"Hola",
		"नमस्कार",
		"こんにちは",
		"Привіт"}
	fmt.Println(greetings)
	subGreetings1 := greetings[:2]
	fmt.Println(subGreetings1)
	subGreetings2 := greetings[1:4]
	fmt.Println(subGreetings2)
	subGreetings3 := greetings[3:]
	fmt.Println(subGreetings3)
}

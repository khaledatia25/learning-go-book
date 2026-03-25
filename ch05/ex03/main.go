package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(str string) string {
		return prefix + " " + str
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}

package main

import "fmt"

func main() {
	var total int
	for i := range 10 {
		total = total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}

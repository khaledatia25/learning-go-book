package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 0, 100)
	for range 100 {
		slice = append(slice, rand.Intn(101))
	}

	for _, v := range slice {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}

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
	fmt.Println(slice)
}

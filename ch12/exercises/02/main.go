package main

import (
	"fmt"
	"math/rand/v2"
)

func numbersGenerator(limit int, ch chan<- int) {
	for range limit {
		ch <- rand.IntN(100)
	}
	close(ch)
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go numbersGenerator(10, ch1)
	go numbersGenerator(10, ch2)

	openChannels := 2
	for openChannels != 0 {
		select {
		case n, ok := <-ch1:
			if !ok {
				ch1 = nil
				openChannels--
				break
			}
			fmt.Println("ch1: ", n)
		case n, ok := <-ch2:
			if !ok {
				ch2 = nil
				openChannels--
				break
			}
			fmt.Println("ch2: ", n)

		}
	}
}

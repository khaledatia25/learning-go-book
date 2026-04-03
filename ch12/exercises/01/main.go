package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func numbersGenerator(limit int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range limit {
		ch <- rand.IntN(1000)
	}
}

func numbersReader(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	var wg, wg2 sync.WaitGroup
	var ch chan int = make(chan int)
	limit := 2
	wg.Add(limit)
	for range 2 {
		go numbersGenerator(10, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	wg2.Add(1)
	go numbersReader(ch, &wg2)
	wg2.Wait()
}

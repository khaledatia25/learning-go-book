package main

import (
	"fmt"
	"math"
	"sync"
)

func mapBuilder() map[int]float64 {
	mp := make(map[int]float64, 100000)
	for i := range 100000 {
		mp[i] = math.Sqrt(float64(i))
	}
	return mp
}

var mapBuilderCached func() map[int]float64 = sync.OnceValue(mapBuilder)

func CachedSqrt(n int) float64 {
	mp := mapBuilderCached()
	return mp[n]
}

func main() {
	for i := 0; i < 100000; i += 1000 {
		fmt.Println(i, ": ", CachedSqrt(i))
	}
}

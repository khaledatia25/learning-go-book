package main

import "fmt"

type Multipliable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Double[T Multipliable](val T) T {
	return val * T(2)
}

func main() {
	var i int8 = 2
	fmt.Println(Double(i))
	var ii int64 = 89
	fmt.Println(Double(ii))

	var f float32 = 3.4
	fmt.Println(f)
	var ff float64 = 55.43
	fmt.Println(Double(ff))
}

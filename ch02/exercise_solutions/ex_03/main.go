package main

import "fmt"

func main() {
	var b byte = 255
	var i int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	i += 1
	bigI += 1

	fmt.Println(b, i, bigI)
}

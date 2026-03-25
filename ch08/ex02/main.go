package main

import "fmt"

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

func Print[T Printable](val T) {
	fmt.Println(val)
}

type PrintableFloat float64

func (pf PrintableFloat) String() string {
	return fmt.Sprintf("my printable float: %f", pf)
}

type PrintableInt int

func (pi PrintableInt) String() string {
	return fmt.Sprintf("my printable int: %d", pi)
}

func main() {
	var pi PrintableInt = 3
	Print(pi)
	var pf PrintableFloat = 5.4
	Print(pf)
}

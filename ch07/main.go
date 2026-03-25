package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return p.FirstName + " " + p.LastName + " " + fmt.Sprint(p.Age)
}

func firstExample() {
	p := Person{
		FirstName: "khaled",
		LastName:  "Atia",
		Age:       24,
	}
	content := p.String()

	fmt.Println(content)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c *Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func secondExample() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}

func main() {
	secondExample()
}

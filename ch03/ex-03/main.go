package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	emp1 := Employee{"Khaled", "Waleed", 1}
	fmt.Println(emp1)
	emp2 := Employee{
		firstName: "Khaled",
		lastName:  "Waleed",
		id:        2,
	}
	fmt.Println(emp2)
	var emp3 Employee
	emp3.firstName = "Khaled"
	emp3.lastName = "Waleed"
	emp3.id = 3
	fmt.Println(emp3)
}

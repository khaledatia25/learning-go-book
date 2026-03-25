package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func main() {
	person1 := MakePerson("Khaled", "Waleed", 21)
	person2 := MakePersonPointer("Kareem", "Hafez", 21)

	fmt.Println(person1)
	fmt.Println(*person2)
}

/*
Task
Create a struct named Person with three fields: FirstName and LastName of type string and Age of type int. Write a function called MakePerson that takes in firstName, lastName, and age and returns a Person. Write a second function MakePersonPointer that takes in firstName, lastName, and age and returns a *Person. Call both from main. Compile your program with go build -gcflags="-m". This both compiles your code and prints out which values escape to the heap. Are you surprised about what escapes?

*/

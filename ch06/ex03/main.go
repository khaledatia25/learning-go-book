package main

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

func main() {
	var people []Person
	for range 10000000 {
		people = append(people, MakePerson("Khaled", "Waleed", 24))
	}
}

/*
Write a program that builds a []Person with 10,000,000 entries (they could all be the same names and ages). See how long it takes to run. Change the value of GOGC and see how that affects the time it takes for the program to complete. Set the environment variable GODEBUG=gctrace=1 to see when garbage collections happen and see how changing GOGC changes the number of garbage collections. What happens if you create the slice with a capacity of 10,000,000?
*/

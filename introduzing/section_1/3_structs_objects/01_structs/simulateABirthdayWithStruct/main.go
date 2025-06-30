package main

/*
	Implement a function that simulates a person's birthday by incrementing their age.
*/

import "fmt"

// Person represents a person with a name and age.
type Person struct {
	Name string
	Age  int
}

// birthday increments the person's age by 1 (simulating a birthday).
func birthday(p *Person) {
	p.Age++ // equivalent to (*p).Age++
}

func main() {
	p := Person{Name: "Ana", Age: 30}

	fmt.Printf("Before birthday: %s is %d years old\n", p.Name, p.Age)

	birthday(&p)

	fmt.Printf("After birthday:  %s is %d years old\n", p.Name, p.Age)
}

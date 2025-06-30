package main

/*
	Check if a name exists in a map of people and their ages.
	If the name exists, print the age. If not, print a message.
*/

import "fmt"

func main() {
	// Create a map with names and ages
	ages := map[string]int{
		"Ana":  23,
		"João": 30,
	}

	// Name to search in the map
	name := "Pedro"

	// Check if the name exists in the map
	if age, ok := ages[name]; ok {
		fmt.Printf("%s is %d years old.\n", name, age)
	} else {
		fmt.Printf("%s is not in the map.\n", name)
	}
}

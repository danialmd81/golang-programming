package main

// Importing fmt
import (
	"fmt"
)

// Calling main
func main() {

	// Declaring some variables
	var name string
	var age int

	// Calling Sscanln() function
	number, error := fmt.Sscanln("Robert 30", &name, &age)

	// Checking if the function returns any error
	if error != nil {
		panic(error)
	}

	// Printing the number of elements present in the specified string and also the elements
	fmt.Printf("Number of Elements: %d, Name: %s, Age: %d", number, name, age)

}

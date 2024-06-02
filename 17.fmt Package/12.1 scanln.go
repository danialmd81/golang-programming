package main

// Importing fmt
import (
	"fmt"
)

func main() {

	// Declaring some variables
	var name string
	var age int

	// Calling Scanln() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Scanln(&name)
	fmt.Scanln(&age)

	// Printing the given texts
	fmt.Printf("%s %d", name, age)

}

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
	var mark float32
	var isAccept bool

	// Calling the Sscan() function which
	// returns the number of elements
	// successfully scanned and error if
	// it persists
	number, error := fmt.Sscan("Robert 25 79.85 true",
		&name, &age, &mark, &isAccept)

	// Below statements get
	// executed if there is any error
	if error != nil {
		panic(error)
	}

	// Printing the number of
	// elements and each elements also
	fmt.Printf("Number of Elements:%d - %s, %d, %f, %t", number, name,
		age, mark, isAccept)

}

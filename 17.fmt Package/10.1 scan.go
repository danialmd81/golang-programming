package main

import (
	"fmt"
)

func main() {

	// Declaring some variables
	var name string
	var age int
	var mark float32
	var isAccept bool

	// Calling Scan() function for scanning and reading the input texts given in standard input
	fmt.Scan(&name)
	fmt.Scan(&age)
	fmt.Scan(&mark)
	fmt.Scan(&isAccept)

	// Printing the given texts
	fmt.Printf("%s %d %.2f %t", name, age, mark, isAccept)

}

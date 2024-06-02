package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Declaring some type of variables
	var name string
	var age int
	var mark float32
	var isAccept bool

	// Calling the NewReader() function to specify some type of texts.variable "scan" contains the scanned texts
	scan := strings.NewReader("Robert 35 95.50 true")

	// Calling the Fscan() function to receive
	// the scanned texts
	number, error := fmt.Fscan(scan, &name, &age, &mark, &isAccept)

	// If the above function returns an error then below statement will be executed
	if error != nil {
		fmt.Fprintf(os.Stderr, "Fscan: %v\n", error)
	}

	// Printing each type of scanned texts and number of items successfully scanned
	fmt.Println("Number of Elements:", number, name, age, mark, isAccept)
}

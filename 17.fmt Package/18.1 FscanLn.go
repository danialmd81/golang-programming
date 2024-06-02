package main

// Importing fmt, io and strings
import (
	"fmt"
	"strings"
)

// Calling main
func main() {

	// Declaring list of strings,integers and float value
	text := "Robert 27 80.53 true"

	// Calling NewReader() function for reading each elements of the list and then it place it into "scan"
	scan := strings.NewReader(text)

	// Declaring different types of variables
	var name string
	var age int
	var mark float32
	var isAccept bool

	// Calling Fscanln() function
	number, error := fmt.Fscanln(scan, &name, &age, &mark, &isAccept)

	// Checking if there is any error
	if error != nil {
		panic(error)
	}

	// Printing the number of items successfully scanned and each elements too
	fmt.Printf("Number of Elements:%d: %s, %d, %f, %t", number, name, age, mark, isAccept)
}

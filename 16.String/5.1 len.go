package main

import (
	"fmt"
	"unicode/utf8"
)

// Main function
func main() {

	// Creating and initializing a string
	// using shorthand declaration

	str := "Welcome to Go Programming Language"
	fmt.Println("string:", str)

	// Finding the length of the string Using len() function
	length1 := len(str)
	// Displaying the length of the string
	fmt.Println("Length 1:", length1)

	//Finding the length of the string Using RuneCountInString() function
	length2 := utf8.RuneCountInString(str)
	// Displaying the length of the string
	fmt.Println("Length 2:", length2)

}

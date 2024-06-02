package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	// Creating and initializing string Using shorthand declaration
	str1 := "Welcome to Go Programming Language"
	str2 := "This is the tutorial of Golang"

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println(str1)
	fmt.Println(str2)

	//Trimming prefix string from the given strings Using TrimPrefix() function
	result1 := strings.TrimPrefix(str1, "Welcome to ")
	result2 := strings.TrimPrefix(str2, "Language")

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println(result1)
	fmt.Println(result2)
}

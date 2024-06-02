package main

import (
	"fmt"
	"strings"
)

func main() {

	// Creating and initializing string Using shorthand declaration
	str1 := "!!Welcome to Go Programming Language **"
	str2 := "@@This is the tutorial of Golang$$"

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)

	// Trimming the given strings Using TrimRight() function
	result1 := strings.TrimRight(str1, "!*")
	result2 := strings.TrimRight(str2, "$")

	// Displaying the resultults
	fmt.Println("\nStrings after trimming:")
	fmt.Println("resultult 1: ", result1)
	fmt.Println("resultult 2:", result2)
}

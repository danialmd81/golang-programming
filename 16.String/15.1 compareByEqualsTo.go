package main

import "fmt"

// Main function
func main() {

	// Creating and initializing strings using shorthand declaration
	str1 := "Go"
	str2 := "GO"
	str3 := "Golang"
	str4 := "Go"

	// Checking the string are equal or not using == operator
	result1 := str1 == str2
	result2 := str2 == str3
	result3 := str3 != str4
	result4 := str1 == str4

	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
	fmt.Println("Result 3: ", result3)
	fmt.Println("Result 4: ", result4)
}

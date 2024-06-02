package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a string
	str := "Go Programming"

	// Accessing the bytes of the given string
	for i := 0; i < len(str); i++ {

		fmt.Printf("Index %d = %c and ASCI code = %v\n", i, str[i], str[i])
	}
}

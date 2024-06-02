package main

import (
	"fmt"
	"strings"
)

// Main function
func main() {

	// Creating and initializing the strings
	str1 := "Welcome, to the, Go Programming, Language"
	str2 := "My name is Robert"
	str3 := "I am a Golang Programmer"

	// Displaying strings
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)
	fmt.Println("String 3: ", str3)

	// Splitting the given strings Using SplitAfter() function
	result1 := strings.SplitAfterN(str1, ",", 3)
	result2 := strings.SplitAfterN(str2, "", 4)
	result3 := strings.SplitAfterN(str3, "!", 1)

	// Displaying the result
	fmt.Printf("\nResult 1 Type:%T \t Result3:%v\n", result1, result1)

	for _, value := range result1 {
		fmt.Println(value)
	}

	fmt.Println("\nResult 2: ", result2)
	fmt.Println("Result 3: ", result3)

}

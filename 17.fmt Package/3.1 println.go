package main

// Importing fmt
import (
	"fmt"
)

// Calling main
func main() {

	// Declaring some string variables
	var name, value = "Go", "Programming Language"

	// Calling Printf() function
	fmt.Println(name, " is a ", value)

	// Declaring some string variables again
	var num1, num2, num3, num4 = 5, 10, 15, 50

	// Calling Printf() function
	fmt.Println(num1, "+", num2, "=", num3)
	fmt.Println(num1, "*", num2, "=", num4)

}

package main

import (
	"fmt"
	"os"
)

// Calling main
func main() {

	// Declaring some const variables
	var num1, num2, num3 = 5, 15, 15

	// Calling Fprint() function which returns "number" as the number of bytes written and
	// "error" as any error encountered
	number, error := fmt.Fprintf(os.Stdout, "%d + %d = %d.\n", num1, num2, num3)

	// Printing the number of bytes written
	fmt.Print(number, " bytes written.\n")

	// Printing if any error encountered
	fmt.Print(error)

}

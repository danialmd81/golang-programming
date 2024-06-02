package main

// Importing fmt
import (
	"fmt"
	"os"
)

func main() {

	// Declaring some string variables
	var name, value = "Go", "Programming Language"

	// Calling Fprintln() function
	number, error := fmt.Fprintln(os.Stdout, name, "is a", value, ".")

	// Printing the number of bytes written
	fmt.Print(number, " bytes written.\n")

	// Printing if any error encountered
	fmt.Print(error)

	fmt.Println("\n------------------")

	// Declaring some const variables
	const str1, str2, str3 = "a", "b", "c"

	// Calling Fprintln() function which returns "number" as the number of bytes written and
	// "error" as any error ancountered
	number, error = fmt.Fprintln(os.Stdout, str1, str2, str3)

	// Printing the number of bytes written
	fmt.Print(number, " bytes written.\n")

	// Printing if any error encountered
	fmt.Print(error)

}

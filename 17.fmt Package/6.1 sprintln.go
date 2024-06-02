package main

// Importing fmt
import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Declaring some string variables
	var name, value = "Go", "Programming Language"

	// Calling Printf() function
	result := fmt.Sprintln(name, " is a ", value)

	// Calling WriteString() function to write the contents of the string "result" to "os.Stdout"
	io.WriteString(os.Stdout, result)

	// Declaring some string variables again
	var num1, num2, num3, num4 = 5, 10, 15, 50

	// Calling Printf() function
	result1 := fmt.Sprintln(num1, "+", num2, "=", num3)
	result2 := fmt.Sprintln(num1, "*", num2, "=", num4)

	// Calling WriteString() function to write the contents of the string "result" to "os.Stdout" again
	io.WriteString(os.Stdout, result1)
	io.WriteString(os.Stdout, result2)

}

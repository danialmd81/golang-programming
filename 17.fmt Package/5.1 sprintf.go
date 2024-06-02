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

	// Calling Sprintf() function
	result := fmt.Sprintf("%s is a %s.\n", name, value)

	// Calling WriteString() function to write the contents of the string "result" to "os.Stdout"
	io.WriteString(os.Stdout, result)

	// Declaring some string variables again
	var num1, num2, num3 = 5, 10, 15

	// Calling Printf() function
	result = fmt.Sprintf("%d + %d = %d", num1, num2, num3)

	// Calling WriteString() function to write the contents of the string "result" to "os.Stdout" again
	io.WriteString(os.Stdout, result)

}

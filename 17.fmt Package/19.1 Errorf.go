package main

// Importing fmt and time
import (
	"fmt"
	"time"
)

// Calling main
func main() {

	// Calling Errorf() function
	err := fmt.Errorf("error occurred at: %v", time.Now())

	// Printing the error
	fmt.Println("An error happened:", err)
}

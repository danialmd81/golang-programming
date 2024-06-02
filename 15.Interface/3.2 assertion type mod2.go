package main

import "fmt"

func display(a interface{}) {

	// Extracting the value of a
	value, ok := a.(string)
	fmt.Printf("Value: %v -- Status is %v\n", value, ok)
}
func main() {

	var result1 = "Go Programming Language"
	display(result1)

	var result2 = 123
	display(result2)

}

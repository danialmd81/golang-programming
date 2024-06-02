package main

import "fmt"

func display(a interface{}) {

	// Extracting the value of a
	value := a.(string)
	fmt.Println("Value: ", value)
}
func main() {

	var result = "Go Programming Language"

	display(result)
}

package main

import "fmt"

func main() {

	// Creating and initializing a variable with a string Using shorthand declaration
	name1 := "Kim"

	// Creating and initializing a variable with a string Using var keyword and
	// Using double-quote and adding escape character
	var name2 string = "Go Programming \n Language"

	// Creating and initializing a variable with a string Using var keyword and
	// Using backticks asin raw literals and adding escape character
	var name3 string = `Go Programming \n Language`

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)
}

package main

import "fmt"

func main() {

	// Passing arguments in anonymous function
	func(name string) {
		fmt.Println(name)
	}("Welcome to go Programming Language")

}

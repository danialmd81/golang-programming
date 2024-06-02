package main

import "fmt"

func main() {

	//Assigning an anonymous function to a variable
	result := func() {
		fmt.Println("Welcome To Go Programming Language ")
	}

	//call result variable as function type
	result()

}

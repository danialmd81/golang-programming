package main

import "fmt"

func main() {

	// Creating array Using shorthand declaration
	array1 := [3]int{1, 2, 3}
	array2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	// Finding the length of the array using len function
	fmt.Println("Length of the array 1 is:", len(array1))
	fmt.Println("Length of the array 2 is:", len(array2))
}

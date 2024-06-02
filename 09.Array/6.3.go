// Go program to illustrate the
// concept of ellipsis in an array
package main

import "fmt"

func main() {

	// Creating an array whose size is determined by the number of elements present in it Using ellipsis
	array := [...]string{"Go", "Java", "C#", "C++", "Perl"}

	fmt.Println("Elements of the array: ", array)

	// Length of the array is determine by Using len() function
	fmt.Println("Length of the array is:", len(array))
}

// Golang program to illustrate the
// concept of a pointer to a struct

package main

import (
	"fmt"
)

// Person Struct
type Person struct {

	// declaring variables
	name string
	id   int64
}

// main function
func main() {

	// creating the instance of the Person struct type
	person1 := Person{"Kim", 112233}

	fmt.Println(person1)

	// referencing the struct person to another variable by using the ampersand operator
	// Here, it is the pointer to the struct
	person2 := &person1

	fmt.Println(person2)

	// changing values of struct p2
	person2.name = "Edvard"
	person2.id = 445566

	fmt.Println("Display two persons after changing")

	// printing updated struct
	fmt.Println(person2)

	// struct person1 values will also change since values of person2 were also changed
	fmt.Println(person1)

}

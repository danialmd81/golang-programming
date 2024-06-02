package main

import "fmt"

// Student Struct
type Student struct {
	FirstName string
	Age       int
}

func main() {

	// Creating variables of Employee structure
	student1 := Student{
		FirstName: "Kim",
		Age:       27,
	}

	student2 := Student{
		FirstName: "Kim",
		Age:       27,
	}

	// Checking if student1 is equal to student2 or not Using == operator
	if student1 == student2 {

		fmt.Println("Variable student1 is equal to variable student2")

	} else {

		fmt.Println("Variable student1 is not equal to variable student2")
	}
}

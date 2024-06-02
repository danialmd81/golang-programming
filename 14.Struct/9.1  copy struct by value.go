package main

import (
	"fmt"
)

// Student struct
type Student struct {

	// declaring variables
	name  string
	marks int64
}

// main function
func main() {

	// creating the instance of the Student struct type
	student1 := Student{"Julia", 90}

	fmt.Println("Student1: ", student1)

	// copying the struct student to another variable by using the assignment operator
	student2 := student1

	fmt.Println("Student2: ", student2)

	// changing values of struct student2 after copying
	student2.name = "Robert"
	student2.marks = 95

	// printing updated struct
	fmt.Println("Display two students after changes")
	fmt.Println("Student1: ", student1)
	fmt.Println("Student2: ", student2)

}

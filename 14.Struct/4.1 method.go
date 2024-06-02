package main

import "fmt"

// Author structure
type author struct {
	name   string
	branch string
	salary int
}

// Method with a receiver of author type
func (a author) show() {

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Salary: ", a.salary)
}

func main() {

	// Initializing the values of the author structure
	result := author{
		name:   "Kim",
		branch: "CSE",
		salary: 29000,
	}

	// Calling the method
	result.show()
}

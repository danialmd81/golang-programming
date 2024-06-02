package main

import "fmt"

// Author structure
type author struct {
	name      string
	branch    string
	particles int
}

// Method with a receiver of author type
func (a *author) show(abranch string) {
	(*a).branch = abranch
}

// Main function
func main() {

	// Initializing the values of the author structure
	result := author{
		name:   "Robert",
		branch: "CSE",
	}

	fmt.Println("Author's name: ", result.name)
	fmt.Println("Branch Name(Before): ", result.branch)

	// Creating a pointer
	p := &result
	//OR
	// result.show("ECE")

	// Calling the show method
	p.show("ECE")
	fmt.Println("Author's name: ", result.name)
	fmt.Println("Branch Name(After): ", result.branch)
}

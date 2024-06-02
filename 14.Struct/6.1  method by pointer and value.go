package main

import "fmt"

type author struct {
	name   string
	branch string
}

// Method with a pointer receiver of author type
func (a *author) show1(abranch string) {
	(*a).branch = abranch
}

// Method with a value receiver of author type
func (a author) show2() {

	a.name = "Julia"
	fmt.Println("Author's name(Before) : ", a.name)
}

// Main function
func main() {

	// Initializing the values of the author structure
	result := author{
		name:   "Kim",
		branch: "CSE",
	}

	fmt.Println("Branch Name(Before): ", result.branch)

	// Calling the show1 method(pointer method) with value
	result.show1("ECE")
	// result.show2()
	fmt.Println("Branch Name(After): ", result.branch)

	// Calling the show2 method (value method) with a pointer
	(&result).show2()
	// (&result).show1("ECE")
	fmt.Println("Author's name(After): ", result.name)
}

package main

import "fmt"

//Author Struct - Creating structure
type Author struct {
	name   string
	branch string
	year   int
}

//HR Struct - Creating nested structure
type HR struct {

	// creating a structure as a field
	details Author
}

func main() {

	// Initializing the fields of the structure
	result := HR{

		details: Author{"Julia", "ECE", 2021},
	}

	// Display the values
	fmt.Println("Details of Author")
	fmt.Println(result)
}

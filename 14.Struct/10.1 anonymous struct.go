package main

import "fmt"

func main() {

	// Creating and initializing the anonymous structure
	student := struct {
		name string
		id   int64
		age  int32
	}{
		name: "Julia",
		id:   12345,
		age:  30,
	}

	// Display the anonymous structure
	fmt.Println(student)
}

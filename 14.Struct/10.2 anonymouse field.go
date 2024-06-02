package main

import "fmt"

type student struct {
	string
	int
}

func main() {

	value := student{"Kim", 123}
	fmt.Println("Student Name:", value.string)
	fmt.Println("Student ID:", value.int)
}

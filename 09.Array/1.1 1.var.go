package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Jim"
	names[1] = "Kim"
	names[2] = "Alex"

	fmt.Println("Elements of Array:")
	fmt.Println("Element1:", names[0])
	fmt.Println("Element2:", names[1])
	fmt.Println("Element3:", names[2])
}

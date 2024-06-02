package main

import "fmt"

func main() {

	var i int = 100
	var s string = "Hello"

	kind(i)
	kind(s)
}

func kind(i interface{}) {
	fmt.Printf("The Type of %v is %T\n", i, i)
}

package main

import "fmt"

func main() {

	var a = 2.35

	var b = 3.76

	var c = complex(a, b)

	fmt.Println(c)
	fmt.Printf("Type of C : %T", c)
}

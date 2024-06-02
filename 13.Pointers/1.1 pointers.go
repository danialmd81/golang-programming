package main

import "fmt"

func main() {

	//Define a General varriable
	var a int = 20

	//Define a Pointer
	var p *int

	//Initialize Pointer 
	p = &a

	fmt.Printf("Address of A Variable:%x\n", &a)
	fmt.Printf("Address stored in P Variable:%x\n", p)
	fmt.Printf("Value of *p Variable:%d\n", *p)
}

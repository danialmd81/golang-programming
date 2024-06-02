package main

import "fmt"

func main() {

	//Define a General varriable
	var a int = 458

	//Define a Pointer
	var p *int

	//Initialize Pointer
	p = &a

	fmt.Println("Value stored in A variable: ", a)
	fmt.Println("Address of A variable: ", &a)

	fmt.Println("Value stored in Pointer P : ", p)
	fmt.Println("Value stored in *P Before Changing: ", *p)

	*p = 500

	fmt.Println("Value stored in *P After Changing: ", a)
}

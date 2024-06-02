package main

import "fmt"

func main() {

	//declare multiple variables for different type in the single diclaration
	var variable1, variable2, variable3 = 2, "Go Programming Language", 67.56

	fmt.Printf("Variable1 Value : %d\n", variable1)
	fmt.Printf("Variable1 Type : %T\n", variable1)

	fmt.Printf("Variable2 Value : %s\n", variable2)
	fmt.Printf("Variable2 Type : %T\n", variable2)

	fmt.Printf("Variable3 Value : %f\n", variable3)
	fmt.Printf("Variable3 Type : %T", variable3)

}

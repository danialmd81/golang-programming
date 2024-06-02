package main

import "fmt"

func main() {

	var a int = 100

	var b int = 100

	if a < b {
		fmt.Println("A Value is less Than b Value")
	} else if a == b {
		fmt.Println("A Value is Equal B Value")
	} else {
		fmt.Println("A Value is Greater Than B Value")
	}

}

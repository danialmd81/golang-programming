package main

import "fmt"

func display(a *int) {
	*a = 500
}
func main() {
	var x = 100
	fmt.Printf("The value of x before function call is: %d\n", x)

	display(&x)
	fmt.Printf("The value of x after function call is: %d\n", x)
}

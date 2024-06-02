package main

import "fmt"

func print(a int) {

	a = 500
}

func main() {

	var x = 100
	fmt.Printf("The Value of X before function call is:%d\n", x)

	print(x)

	fmt.Printf("The Value of X After function call is:%d\n", x)
}

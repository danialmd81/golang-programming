package main

import "fmt"

func swap(x, y int) int {
	var temp int

	temp = x
	x = y
	y = temp
	return temp
}

func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, A = %d and B = %d\n", a, b)

	swap(a, b)

	fmt.Printf("After swap, A = %d and B = %d", a, b)
}

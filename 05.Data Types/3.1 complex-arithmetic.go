package main

import "fmt"

func main() {

	var a = 3 + 5i
	var b = 2 + 3i

	fmt.Println("a =", a, "\nb =", b)

	var result1 = a + b

	var result2 = b - a

	var result3 = a * b

	var result4 = a / b

	fmt.Println("a + b =", result1, "\nb - a =", result2, "\na * b =", result3, "\na / b =", result4)
}

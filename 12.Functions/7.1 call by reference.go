package main

import "fmt"

func swap(x *int, y *int) {

	var temp int

	temp = *x

	*x = *y //باید حتما ستاره ها باشن
	*y = temp
}

func main() {

	var a int = 100
	var b int = 200

	fmt.Printf("Before Swap, A = %d and B = %d\n", a, b)

	swap(&a, &b)

	fmt.Printf("After Swap, A = %d and B = %d", a, b)
}

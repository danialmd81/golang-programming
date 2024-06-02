package main

import "fmt"

func main() {

	var x int = 0

lable1:

	for x < 8 {

		if x == 5 {

			x = x + 1
			goto lable1
		}

		fmt.Printf("Value of X : %d\n", x)

		x++
	}
}

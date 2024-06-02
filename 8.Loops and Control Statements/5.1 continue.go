package main

import "fmt"

func main() {

	var x int = 0

	for x < 8 {

		if x == 5 {

			x = x + 2
			continue //میره اول حلقه نه ادامه کد
		}

		fmt.Printf("Value is %d\n", x)
		x++
	}
}

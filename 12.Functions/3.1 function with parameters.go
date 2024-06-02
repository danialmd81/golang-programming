package main

import "fmt"

func add(x int, y int) {

	total := x + y

	fmt.Printf("Total of X + Y = %d", total)
}

func main() {

	add(20, 30)
}

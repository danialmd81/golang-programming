package main

import "fmt"

func add(x int, y int) int {
	
    total := x + y

	return total
}

func main() {
	
    sum := add(20, 30)

	fmt.Printf("Sum of X + Y = %d", sum)
}

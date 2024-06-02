package main

import "fmt"

func main() {

	array := [6]int{10, 20, 30, 40, 50, 60}

	slice := array[0:4]

	fmt.Println("First Array:", array)

	fmt.Println("First Slice:", slice)

	slice[0] = 100
	slice[1] = 200
	slice[2] = 300

	fmt.Println("\nSecond Array:", array)
	fmt.Println("Second Slice:", slice)

}

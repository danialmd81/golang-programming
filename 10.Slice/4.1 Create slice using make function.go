package main

import "fmt"

func main() {

	var slice1 = make([]int, 4, 7)

	fmt.Printf("Slice1 = %v,\tLength = %d,\tCapacity = %d\n", slice1, len(slice1), cap(slice1))

	var slice2 = make([]int, 7)

	fmt.Printf("Slice2 = %v,\tLength = %d,\tCapacity = %d\n", slice2, len(slice2), cap(slice2))
}

package main

import "fmt"

func main() {

	array := [5]int{1, 2, 3, 4, 5}

	var slice1 = array[1:2]

	slice2 := array[0:]

	slice3 := array[:2]

	slice4 := array[:]

	fmt.Println("Array:", array)

	fmt.Println("Slice1:", slice1)

	fmt.Println("Slice2:", slice2)

	fmt.Println("Slice3:", slice3)

	fmt.Println("Slice4:", slice4)
}

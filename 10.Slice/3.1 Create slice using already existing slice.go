package main

import "fmt"

func main() {

	mainSlice := []int{10, 20, 30, 40, 50, 60, 70}

	var slice1 = mainSlice[1:5]

	slice2 := mainSlice[0:]

	slice3 := mainSlice[:6]

	slice4 := mainSlice[:]

	slice5 := mainSlice[2:4]

	fmt.Println("Main Slice:", mainSlice)

	fmt.Println("Slice1:", slice1)

	fmt.Println("Slice2:", slice2)

	fmt.Println("Slice3:", slice3)

	fmt.Println("Slice4:", slice4)

	fmt.Println("Slice5:", slice5)
}

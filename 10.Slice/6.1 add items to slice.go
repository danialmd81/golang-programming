package main

import "fmt"

func main() {

	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20

	fmt.Println("Slice A:", a)

	fmt.Printf("Length:%d\tCapacity:%d\n", len(a), cap(a))

	a = append(a, 30, 40, 50, 60, 70, 80, 90)

	fmt.Println("Slice A after Appending Data:", a)

	fmt.Printf("Length:%d\tCapacity:%d\n", len(a), cap(a))
}

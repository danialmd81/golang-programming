package main

import "fmt"

func main() {

	a := []int{10, 20, 30}

	fmt.Println("Slice A:", a)

	fmt.Printf("Slice A Length:%d and Capacity:%d\n", len(a), cap(a))

	b := make([]int, 5, 10)

	copy(b, a)

	fmt.Println("Slice B:", b)

	fmt.Printf("Slice B Length:%d and Capacity:%d", len(b), cap(b))
}

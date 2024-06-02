package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70}

	for index, value := range slice {
		fmt.Printf("Index = %d and Value = %d\n", index, value)
	}

}

package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70}

	for _, value := range slice {
		fmt.Printf("The Value = %d\n", value)
	}

}

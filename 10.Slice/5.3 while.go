package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70}

	i := 0

	for range slice {
		fmt.Println(slice[i])
		i++
	}

}

package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

}

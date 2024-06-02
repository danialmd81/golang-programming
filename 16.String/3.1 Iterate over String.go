package main

import "fmt"

func main() {
	value := "Go programming Language"

	for index, ch := range value {
		fmt.Printf("Value of index %d is %c\n", index, ch)

	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {

	// Comparing string using Compare function
	fmt.Println(strings.Compare("hello", "world"))

	fmt.Println(strings.Compare("Golang", "Golang"))

	fmt.Println(strings.Compare("golang", "Golang"))

}

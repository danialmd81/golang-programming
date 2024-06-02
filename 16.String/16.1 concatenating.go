package main

import (
	"fmt"
	"strings"
)

func main() {
	//Case 1 s contains sep. Will output slice of length 3
	result := strings.Join([]string{"ab", "cd", "ef"}, "-")
	fmt.Println(result)

	//Case 2 slice is empty. It will output a empty string
	result = strings.Join([]string{}, "-")
	fmt.Println(result)

	//Case 3 sep is empty. It will output a string combined from the slice of strings
	result = strings.Join([]string{"gh", "ij", "kl"}, "")
	fmt.Println(result)
}

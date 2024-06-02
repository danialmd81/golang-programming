package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Australia", "Aus")) // Any part of string
	fmt.Println(strings.Contains("Germany", "GER"))   // Case sensitive
	fmt.Println(strings.Contains(" ", " "))           // space also consider as string
	fmt.Println(strings.Contains("123", "1"))

}

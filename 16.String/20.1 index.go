package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("Australia", "Aus"))
	fmt.Println(strings.Index("Australia", "aus"))
	fmt.Println(strings.Index("Australia", "A"))
	fmt.Println(strings.Index("Australia", "a"))
}

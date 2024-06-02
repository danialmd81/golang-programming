package main

import (
	"fmt"
	"strings"
)

func main() {

	message := "Welcome to Go Programming Language"
	messageArray := strings.Fields(message)

	for _, v := range messageArray {
		fmt.Println(v)
	}

}

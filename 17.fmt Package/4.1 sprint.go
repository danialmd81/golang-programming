package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	var name, value = "Go", "Programming Language"

	result := fmt.Sprint(name, " is a ", value)

	io.WriteString(os.Stdout, result)
}

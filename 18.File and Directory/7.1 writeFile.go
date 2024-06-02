package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Create("test.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	len, err := file.WriteString("Hello and Welcome to\n Go Programming Language.")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("File Name: %s\n", file.Name())
	fmt.Printf("Length: %d bytes", len)
}

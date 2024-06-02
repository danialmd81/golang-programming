package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//Create a file
	file, err := os.Create("temp.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Write something to the file
	file.WriteString("some sample text" + "\n")

	//Gets stats of the file
	stats, err := os.Stat("temp.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Prints stats of the file
	fmt.Printf("Permission: %s\n", stats.Mode())
	fmt.Printf("Name: %s\n", stats.Name())
	fmt.Printf("Size: %d\n", stats.Size())
	fmt.Printf("Modification Time: %s\n", stats.ModTime())

	file.Close()
}

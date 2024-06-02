package main

import (
	"log"
	"os"
)

func main() {
	//Create a directory
	err := os.Mkdir("temp", 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename("temp", "newTemp")
	if err != nil {
		log.Fatal(err)
	}
}

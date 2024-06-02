package main

import (
	"log"
	"os"
)

func main() {

	// Remove all the directories and files Using RemoveAll() function
	err := os.RemoveAll("C:\\Users\\Hamidreza\\go\\src\\sample")
	if err != nil {
		log.Println(err)
	}
}

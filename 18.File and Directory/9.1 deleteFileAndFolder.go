package main

import (
	"log"
	"os"
)

func main() {

	// Removing file from the current Path Using Remove() function
	fileError := os.Remove("test.txt")
	if fileError != nil {
		log.Println(fileError)
	}

	// Removing folder from the current Path Using Remove() function
	folderError := os.Remove("temp")
	if folderError != nil {
		log.Println(folderError)
	}

}

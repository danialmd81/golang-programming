package main

import (
	"log"
	"os"
)

func main() {
	folderInfo, error := os.Stat("temp")
	if os.IsNotExist(error) {
		log.Fatal("\nFolder does not exist.")
	}
	log.Println(folderInfo)
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileinfo, error := os.Stat("temp.txt")
	if os.IsNotExist(error) {
		log.Fatal("\nFile does not exist.")
	}
	log.Println(fileinfo)
	fmt.Println(error.Error())
}

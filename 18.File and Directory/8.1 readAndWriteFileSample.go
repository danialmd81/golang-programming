package main

// importing the requires packages
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// CreateFile Function
func CreateFile(filename, text string) {

	// Creating the file using Create() function
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// closing the running file after the main function has completed execution
	defer file.Close()

	// writing data to the file using WriteString() function
	len, err := file.WriteString(text)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

// ReadFile function
func ReadFile(filename string) {

	fmt.Printf("\n\nReading a file in Go lang\n")

	// file is read using ReadFile() method of ioutil package
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", filename)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

// main function
func main() {

	// user input for filename
	fmt.Println("Enter filename: ")
	var filename string
	fmt.Scanln(&filename)

	// user input for file content
	fmt.Println("Enter text: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')

	// file is created and read
	CreateFile(filename, input)
	ReadFile(filename)
}

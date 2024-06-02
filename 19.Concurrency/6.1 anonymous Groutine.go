package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello to Main Function")

	go func() {

		fmt.Println("Welcome to Goroutine ......")

	}()

	time.Sleep(1 * time.Second)

	fmt.Println("Goodbye")

}

package main

import (
	"fmt"
	"time"
)

func main() {

	go hello1()

	fmt.Println("Hello")

	time.Sleep(1 * time.Second)

	fmt.Println("Goodbye")
}

func hello1() {

	go hello2()
	fmt.Println("In Goroutine1")
}

func hello2() {

	fmt.Println("In Goroutine2")
}

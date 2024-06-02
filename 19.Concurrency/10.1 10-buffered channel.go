package main

import "fmt"

func main() {

	ch := make(chan string, 1)

	ch <- "Golang"

	message := <-ch

	fmt.Println(message)
}

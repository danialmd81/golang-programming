package main

import "fmt"

func main() {

	ch := make(chan int, 5)

	ch <- 10

	fmt.Printf("Len:%d\t", len(ch))
	fmt.Printf("Cap:%d\t\n", cap(ch))

	ch <- 20

	fmt.Printf("Len:%d\t", len(ch))
	fmt.Printf("Cap:%d\t\n", cap(ch))

	ch <- 30

	fmt.Printf("Len:%d\t", len(ch))
	fmt.Printf("Cap:%d\t\n", cap(ch))
}

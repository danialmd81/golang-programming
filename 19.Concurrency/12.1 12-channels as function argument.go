package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 3)

	go process(ch)

	time.Sleep(1 * time.Second)
}

func process(ch chan int) {

	ch <- 10

	value := <-ch

	fmt.Println(value)
}

package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 3)

	ch <- 10

	go process(ch)

	fmt.Println(<-ch)

}

func process(ch <-chan int) {

	value := <-ch

	fmt.Println(value)
}

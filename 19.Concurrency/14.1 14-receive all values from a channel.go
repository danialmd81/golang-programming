package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 3)

	ch <- 2
	ch <- 2
	ch <- 2

	close(ch)

	sum(ch)

	time.Sleep(time.Millisecond * 200)

}

func sum(ch chan int) {

	sum := 0

	for value := range ch {
		sum += value
	}

	fmt.Printf("Sum Total:%d\n", sum)
}

package main

import "fmt"

func main() {

	ch := make(chan string)

	go send(ch)

	for {

		result, ok := <-ch
		if !ok {

			fmt.Println("Channel Close", ok)
			break
		}

		fmt.Println("Channel Open", result, ok)

	}
}

func send(ch chan string) {

	for i := 0; i < 4; i++ {

		ch <- "Go Programming"
	}

	close(ch)
}

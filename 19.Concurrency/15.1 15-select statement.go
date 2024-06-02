package main

import "fmt"

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	// go bye(ch2)
	// go hello(ch1)
	go hello(ch1)
	go bye(ch2)


	select {

	case value1 := <-ch1:
		fmt.Println(value1)

	case value2 := <-ch2:
		fmt.Println(value2)
	}

}

func hello(ch chan string) {

	ch <- "Hello"
}

func bye(ch chan string) {

	ch <- "Goodbye"
}

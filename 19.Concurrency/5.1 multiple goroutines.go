package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Main Goroutine Start ........")

	go name()

	go id()

	time.Sleep(2000 * time.Millisecond)

	fmt.Println("\nMain Goroutine End ........")
}

func name() {

	array1 := [4]string{"Kim", "Jessica", "Julia", "Tom"}
	for i := 0; i < 4; i++ {

		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%s\n", array1[i])
	}
}

func id() {

	array2 := [4]int{300, 301, 302, 303}
	for i := 0; i < 4; i++ {

		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%d\n", array2[i])
	}
}

package main

import "fmt"

func main() {

	fmt.Printf("Rectangle Area by 10 value = %d\n", func(w int) int {
		return w * w
	}(10))
}

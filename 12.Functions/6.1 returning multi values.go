package main

import "fmt"

func rectangle(w int, h int) (int, int) {

	perimeter := 2 * (w + h)

	area := w * h

	return area, perimeter
}

func main() {

	var a, b int

	a, b = rectangle(20, 30)

	fmt.Println("Area:", a)

	fmt.Println("Perimeter:", b)
}

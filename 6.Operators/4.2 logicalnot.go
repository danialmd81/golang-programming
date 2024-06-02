package main

import "fmt"

func main() {

	//Logical Not ( ! )
	var x, y, z = 100, 200, 300

	fmt.Println(!(x == y && x > z))
}

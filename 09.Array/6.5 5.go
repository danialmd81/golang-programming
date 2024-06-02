package main

import "fmt"

func main() {

	// Arrays
	array1 := [5]int{1, 2, 3, 4, 5}
	array2 := [...]int{1, 2, 3, 4, 5}
	array3 := [5]int{6, 2, 3, 7, 9}

	// Comparing arrays using == operator
	fmt.Println(array1 == array2)
	fmt.Println(array2 == array3)
	fmt.Println(array1 == array3)

}

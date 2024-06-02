package main

import "fmt"

func main() {

	array1 := [3]string{"Kim", "Jim", "Bill"}

	array2 := array1

	array3 := &array1

	fmt.Printf("Array1:%v\n", array1)
	fmt.Printf("Array2:%v\n", array2)
	fmt.Printf("Array3:%v\n", *array3)

	array1[0] = "Alex"

	fmt.Printf("Array1:%v\n", array1)
	fmt.Printf("Array2:%v\n", array2)

	fmt.Printf("Array3:%v", *array3)

}

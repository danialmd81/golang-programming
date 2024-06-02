package main

import "fmt"

func main() {

	names := [5]string{"Kim", "Jim", "Bill", "Robert", "David"}

	fmt.Printf("Name:%v\n", names)
	//Name:[Kim Jim Bill Robert David]
	fmt.Println(names[:]) //from index 0 until end of array
	// [Kim Jim Bill Robert David]
	fmt.Println(names[:3]) //from index 0 until  index 3 of array
	// [Kim Jim Bill]
	fmt.Println(names[2:]) //from index 2 until end of array
	// [Bill Robert David]
	fmt.Println(names[1:4]) //from index 1 until index 4 of array
	// [Jim Bill Robert]
}

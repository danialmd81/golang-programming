package main

import "fmt"

func main() {

	myMap := map[int]string{
		10: "Kim",
		20: "Jim",
		30: "Alex",
	}

	for index, value := range myMap {
		fmt.Printf("index:%d\tValue:%s\n", index, value)
	}

	for i := range myMap {
		delete(myMap, i)
	}

	fmt.Println("My Map After Deleting:")
	fmt.Println(myMap)
}

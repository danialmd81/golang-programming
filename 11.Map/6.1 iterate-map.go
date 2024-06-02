package main

import "fmt"

func main() {

	// Creating and initializing a map
	myMap := map[int]string{10: "Kim", 20: "Edvard", 30: "Robert", 40: "Alex", 50: "David"}

	//display id and name for each element
	for id, name := range myMap {
		fmt.Printf("ID:%d\tName:%s\n", id, name)
	}
}

package main

import "fmt"

func main() {

	//Creating empty map
	var emptyMap = map[int]string{}
	fmt.Println(emptyMap)

	//creating and initializing a simple map
	var myMap = map[int]string{
		1: "kim",
		2: "Jim",
		3: "Robert",
		4: "Jim",
		5: "Edvard",
	}
	fmt.Println(myMap)

}

package main

import "fmt"

func main() {

	// Creating and initializing a map
	myMap := map[int]string{
		10: "Kim",
		20: "Jim",
		30: "Robert",
		40: "Edvard",
		50: "Sofia",
	}

	fmt.Println("Original map: ", myMap)

	// Adding new key-value pairs in the map
	myMap[60] = "Jack"
	myMap[70] = "Julia"
	fmt.Println("Map after adding new key-value pair:\n", myMap)

	// Updating values of the map
	myMap[20] = "Diego"
	fmt.Println("\nMap after updating values of the map:\n", myMap)

	// Deleting values of the map
	delete(myMap, 50)
	fmt.Println("\nMap after deleting values of the map:\n", myMap)
}

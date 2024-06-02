package main

import "fmt"

/*type address struct {
	state string
	city string
	zipCode int
}*/

type address struct {
	state, city string
	zipCode     int
}

func main() {

	//Declaring a variable of a `struct` type
	var a address
	fmt.Println(a)

	//Declaring and initializing a struct using a struct literal
	a1 := address{" Texas ", " Dallas ", 12345}
	fmt.Println("Address1: ", a1)

	// Naming fields while initializing a struct
	a2 := address{state: " California", city: " San Diego", zipCode: 67890}
	fmt.Println("Address2: ", a2)

	// Uninitialized fields are set to their corresponding zero-value
	a3 := address{state: " Texas "}
	fmt.Println("Address3: ", a3)

	//Struct Instantiation using new keyword

	a4 := new(address)
	a4.state = "Illinois"
	a4.city = "Chicago"
	a4.zipCode = 24680
	fmt.Println("Address4:", a4)

	//struct initialization using pointer operator
	var a5 = &address{}
	a5.state = "Colorado"
	a5.city = "Denver"
	a5.zipCode = 13579
	fmt.Println("Address5:", a5)

}

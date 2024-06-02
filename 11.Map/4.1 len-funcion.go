package main

import "fmt"

func main() {

	// Creating a map Using make() function
	var employee = make(map[string]int)

	employee["Kin"] = 10
	employee["Jim"] = 20
	employee["Robert"] = 30

	employeeList := make(map[string]int)
	fmt.Println(employee)
	fmt.Println(len(employee))
	fmt.Println(employeeList)
	fmt.Println(len(employeeList))

}

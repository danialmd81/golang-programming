// Go program to illustrate how to use type switch
package main

import "fmt"

func display(a interface{}) {

	// Using type switch
	switch a.(type) {

	case int:
		fmt.Printf("Type:%T -- Value:%v\n", a, a.(int))
	case string:
		fmt.Printf("Type:%T -- Value:%v\n", a, a.(string))
	case float64:
		fmt.Printf("Type:%T -- Value:%v\n", a, a.(float64))
	default:
		fmt.Println("Type not found")
	}
}

// Main method
func main() {

	display(123)
	display("Go Programming Language")
	display(100.55)
	display(true)
}

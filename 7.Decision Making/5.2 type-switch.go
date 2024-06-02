package main

import "fmt"

func main() {

	var value interface{} = "Go Programming Language"

	switch a := value.(type) {
	case int64:
		fmt.Println("Type is Integer : ", a)
	case float64:
		fmt.Println("Type is Float : ", a)
	case string:
		fmt.Println("Type is String : ", a)
	default:
		fmt.Println("Type is Unknown")

	}

}

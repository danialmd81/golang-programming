package main

import (
	"fmt"
	"reflect"
)

func main() {

	//decalte string data type
	str := "Go Programming Language"

	fmt.Printf("The Value of str : %s\n", str)

	fmt.Printf("The Type of str : %T\n", str)

	fmt.Printf("The Length of str : %d\n", len(str))

	fmt.Println("The Value of str :", str)

	fmt.Println("The Type of str : ", reflect.TypeOf(str))

	fmt.Println("The Length of str :", len(str))
}

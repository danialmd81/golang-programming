package main

import "fmt"

//Rectangle Interface
type Rectangle interface {

	// define Methods for interfce
	Perimeter() float64
	Area() float64
}

// Creating an struct that it wants to implement defined interface methods

//Calculate Struct
type Calculate struct {
	length float64
	width  float64
}

// Implementing methods of the Rectangle interface

//Perimeter Method
func (c Calculate) Perimeter() float64 {

	return (c.width + c.length) * 2
}

//Area Method
func (c Calculate) Area() float64 {

	return c.width * c.length
}

func main() {

	// Accessing elements of the Rectangle interface
	//first create an interface variable from Rectangle and then initialize it by Calculate struct

	var r Rectangle
	r = Calculate{20, 10}

	//invoke rectangle interface methods and display them
	fmt.Println("Perimeter of Rectangle:", r.Perimeter())
	fmt.Println("Area of Rectangle:", r.Area())
}

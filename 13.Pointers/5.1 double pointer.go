package main

import "fmt"

func main() {

	var x int = 100

	var pt1 *int = &x

	var pt2 **int = &pt1

	fmt.Println("The Value of X = ", x)

	fmt.Println("The Value of PT1 = ", pt1)

	fmt.Println("The Value of PT2 = ", pt2)

	fmt.Println("Value at the Address of PT2 or *pt2 = ", *pt2)

	fmt.Println("Value at the Address of PT2 or **PT2 = ", **pt2)

}

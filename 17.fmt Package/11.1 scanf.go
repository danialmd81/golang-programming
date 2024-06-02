package main

import (
	"fmt"
)

func main() {

	var name string
	var age int
	var mark float32
	var isAccept bool

	fmt.Scanf("%s", &name)
	fmt.Scanf("%d", &age)
	fmt.Scanf("%f", &mark)
	fmt.Scanf("%t", &isAccept)

	fmt.Printf("%s %d %f %t", name, age, mark, isAccept)
}

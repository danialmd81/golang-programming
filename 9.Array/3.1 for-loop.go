package main

import "fmt"

func main() {

	intArray := [5]int{1, 2, 3, 4, 5}

	// for i := 0; i < len(intArray); i++ {
	// 	fmt.Println(intArray[i])
	// }

	for index, value := range intArray {
		fmt.Println(index, "=>", value)
	}

	// for _, value := range intArray {
	// 	fmt.Println(value)
	// }

	// j := 0
	// for range intArray {
	// 	fmt.Println(intArray[j])
	// 	j++
	// }
}

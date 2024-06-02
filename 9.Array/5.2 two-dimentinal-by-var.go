package main

import "fmt"

func main() {

	var array2 [2][2]int

	array2[0][0] = 10
	array2[0][1] = 20

	array2[1][0] = 30
	array2[1][1] = 40

	fmt.Println("Elements of Array2:")

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d\t", array2[i][j])
		}

		fmt.Println("")
	}
}

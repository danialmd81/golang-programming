package main

import "fmt"

func main() {

	var today int = 2
	// fmt.Scanf("%d", &today)
	switch {
	case today == 1:
		fmt.Println("Today is Monday")
	case today == 2:
		fmt.Println("Today is Tuesady")
	case today == 3:
		fmt.Println("Todat is Wednesday")
	case today == 4:
		fmt.Println("Today is Thursday")
	case today == 5:
		fmt.Println("Today is Friday")
	case today == 6:
		fmt.Println("Todat is Wednesday")
	case today == 7:
		fmt.Println("Today is Thursday")
	default:
		fmt.Println("Value for Today is invalid")
	}
}

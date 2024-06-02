package main

import "fmt"

func main() {

	names := [4]string{"Jim", "Kim", "Alex", "David"}

	fmt.Println("Elements of Array:")

	for i := 0; i < 4; i++ {
		fmt.Println(names[i])
	}
}

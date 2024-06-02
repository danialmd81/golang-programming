package main

import "fmt"

func main() {

	var slice1 = []string{"Java", "C#", "Go"}
	var slice2 = []string{"Perl", "PHP"}

	slice2 = append(slice2, slice1...)

	fmt.Println("Slice1:", slice1)

	fmt.Println("Slice2:", slice2)
}

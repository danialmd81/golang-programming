package main

import "fmt"

func main() {

	array1 := [3][3]string{
		{"Go", "Java", "C#"},
		{"C", "Scala", "Perl"},
		{"Python", "C++", "Ruby"},
	}

	fmt.Println("Elements of Array1:")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s\t", array1[i][j])
		}
		fmt.Println("")
	}
}

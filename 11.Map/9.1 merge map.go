package main

import "fmt"

func main() {

	first := map[int]string{
		10: "Kim",
		20: "Jim",
		30: "Alex",
	}

	second := map[int]string{
		40: "Robert",
		20: "Jim",
		50: "Julia",
		60: "Diego",
	}

	for i, j := range second {
		first[i] = j
	}

	for i, j := range first {
		fmt.Printf("index:%d\tvalue:%s\n", i, j)
	}
}

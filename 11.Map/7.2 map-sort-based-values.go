package main

import (
	"fmt"
	"sort"
)

func main() {

	unsortedMap := map[int]string{
		30: "kim",
		50: "Jim",
		40: "Alex",
		20: "Diego",
		10: "Jack",
	}

	values := make([]string, 0, len(unsortedMap))

	for _, value := range unsortedMap {
		values = append(values, value)
	}

	sort.Strings(values)

	for _, value := range values {
		fmt.Println(value)
	}
}

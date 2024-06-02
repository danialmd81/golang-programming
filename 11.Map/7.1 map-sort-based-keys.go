package main

import (
	"fmt"
	"sort"
)

func main() {

	//first create a map and initilize it
	unSortedMap := map[int]string{
		30: "Kim",
		50: "Jim",
		40: "Alex",
		20: "Diego",
		10: "Jack",
	}

	// second create a slice
	keys := make([]int, 0, len(unSortedMap))

	//appenap in slice
	for i := range unSortedMap {
		keys = append(keys, i)
	}

	// sort slice
	sort.Ints(keys)

	//iterate the slice
	for _, value := range keys {
		fmt.Println(value, unSortedMap[value])
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {

	square := func(x float64) float64 {

		return math.Sqrt(x)
	}

	fmt.Println(square(9))
}

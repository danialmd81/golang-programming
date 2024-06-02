package main

import (
	"fmt"
	"strings"
)

func main() {
	res := strings.Replace("abcd-abcd", "ab", "**", 1)
	fmt.Println(res)

	res = strings.Replace("abcd-abcd", "ab", "**", -1)
	fmt.Println(res)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name     string
	Age      int64
	Location string
}

func main() {

	j := []byte(`{"name":"Kim","age":20,"location":"Texas"}`)

	var p Person

	err := json.Unmarshal(j, &p)

	if err != nil {
		log.Fatal("Unable to Decode The Json")
	}

	fmt.Printf("Name:%s\nAge:%d\nLocation:%s\n", p.Name, p.Age, p.Location)
}

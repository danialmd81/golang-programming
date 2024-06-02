package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name     string `json:"firstName"`
	Age      int64
	Location string
}

func main() {

	person := Person{

		"Kim", 20, "Texas",
	}

	personArray, err := json.Marshal(person)

	if err != nil {
		log.Fatal("Unable to Encode")
	}

	fmt.Println(string(personArray))
}

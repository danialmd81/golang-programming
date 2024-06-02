package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Person struct {
	Name     string
	Age      int64
	Location string
}

func main() {

	jsonStream := []byte(`{"Name":"Kim","Age":20,"Location":"Texas"}{"Name":"Jessica","Age":30,"Location":"California"}`)

	reader := strings.NewReader(string(jsonStream))

	writer := os.Stdout

	decoder := json.NewDecoder(reader)

	encoder := json.NewEncoder(writer)

	for {

		var v map[string]interface{}

		if err := decoder.Decode(&v); err != nil {
			log.Println(err)
			return
		}

		for k := range v {
			if k == "Age" {
				delete(v, k)
			}
		}

		if err := encoder.Encode(&v); err != nil {
			log.Println(err)
			return
		}
	}
}

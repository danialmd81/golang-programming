package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"here/web"
)

type Configuration struct {
	ServerPort string `json:"serverport"`
}

func main() {

	file, err := os.Open("config.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	conf := new(Configuration)

	err = json.NewDecoder(file).Decode(conf)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Run Server ........")

	web.RunWeb(conf.ServerPort)
}

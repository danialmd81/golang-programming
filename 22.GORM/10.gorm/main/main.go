package main

import (
	"encoding/json"
	"fmt"
	"gorm/dbtools"
	"gorm/model"
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	DataSourceName string `json:"dataSourceName"`
}

func main() {

	file, err := os.Open("configuration/config.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	conf := new(Configuration)

	json.NewDecoder(file).Decode(conf)

	dbtools.DBInitializer(conf.DataSourceName)

	age := 20

	ageWhereClause := "Age=" + strconv.Itoa(age)

	rows := dbtools.MultipleDelete(model.Student{}, ageWhereClause)

	fmt.Println("Rows Affected:", rows)

}

package main

import (
	"encoding/json"
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
	//-------------------------------------------------------------------------------------------
	// dbtools.CreateTable(&model.Student{})
	//-------------------------------------------------------------------------------------------
	// student := model.Student{
	// 	ID:   3,
	// 	Name: "Mani",
	// 	Age:  34,
	// }
	// dbtools.Save(student)
	// println("New Student has inserted in database")
	//-------------------------------------------------------------------------------------------
	// var students []model.Student
	// dbtools.Select(students)
	//-------------------------------------------------------------------------------------------
	// student := model.Student{
	// 	ID: 1,
	// }
	// data := map[string]interface{}{
	// 	"Name": "Jhon",
	// 	"Age":  40,
	// }
	// row := dbtools.SingleUpdate(&student, data)
	// println("Row Affected:", row)
	//-------------------------------------------------------------------------------------------
	// name := "Kim"
	// nameWgormClause := "Name='" + name + "'"
	// data := map[string]interface{}{
	// 	"Name": "David",
	// 	"Age":  50,
	// }
	// rows := dbtools.MultipleUpdate(model.Student{}, nameWgormClause, data)
	// println("Rows Affected:", rows)
	//-------------------------------------------------------------------------------------------
	// age := 34
	// ageWgormClause := "Age=" + strconv.Itoa(age)
	// data := map[string]interface{}{
	// 	"Name": "David",
	// 	"Age":  50,
	// }
	// rows := dbtools.MultipleUpdate(model.Student{}, ageWgormClause, data)
	// println("Rows Affected:", rows)
	//-------------------------------------------------------------------------------------------
	// student := model.Student{
	// 	ID: 2,
	// }
	// row := dbtools.SingleDelete(&student)
	// println("Row Affected:", row)
	//-------------------------------------------------------------------------------------------
	// name := "David"
	// nameWgormClause := "Name='" + name + "'"
	// rows := dbtools.MultipleDelete(model.Student{}, nameWgormClause)
	// println("Rows Affected:", rows)
	//-------------------------------------------------------------------------------------------
	age := 40
	ageWgormClause := "Age=" + strconv.Itoa(age)
	rows := dbtools.MultipleDelete(model.Student{}, ageWgormClause)
	println("Rows Affected:", rows)

}

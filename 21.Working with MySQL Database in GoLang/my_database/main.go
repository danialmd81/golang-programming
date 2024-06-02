package main

import (
	"encoding/json"
	"log"
	"my_database/dbtools"
	"os"
)

type Configuration struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}

func main() {

	file, err := os.Open("./configuration/config.json")

	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	conf := new(Configuration)

	json.NewDecoder(file).Decode(conf)

	dbtools.DBInitializer(conf.DriverName, conf.DataSourceName)

	//-------------------------------------------------------------------------------------------
	// students := dbtools.SelectAllStudents()
	// for _, student := range students {
	// 	println("ID:", student.ID, "Name:", student.Name, "Age:", student.Age)
	// }
	//-------------------------------------------------------------------------------------------
	// student := dbtools.SelectStudentBasedID(5)
	// println("ID:", student.ID, "Name:", student.Name, "Age:", student.Age)
	//-------------------------------------------------------------------------------------------
	// student := model.Student{
	// 	ID:   9,
	// 	Name: "Raj",
	// 	Age:  25,
	// }
	// LastInsertId := dbtools.Save(student)
	// println("Last Inserted ID:", LastInsertId)
	//-------------------------------------------------------------------------------------------
	// student := model.Student{
	// 	ID:   4,
	// 	Name: "Kim",
	// 	Age:  25,
	// }
	// rowsAffected := dbtools.Update(student)
	// println("RowsAffected:", rowsAffected)
	//-------------------------------------------------------------------------------------------
	rowsAffected := dbtools.Delete(5)
	println("Rows Affetced", rowsAffected)
}

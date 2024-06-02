package dbtools

import (
	"communication/model"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBInitializer struct {
	db *sql.DB
}

func Connect(driverName string, dataSourceName string) (*DBInitializer, error) {

	dbconnection, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &DBInitializer{
		db: dbconnection,
	}, nil
}

func (dbInitializer *DBInitializer) SelectStudentBasedID(id int64) (model.Student, error) {

	row := dbInitializer.db.QueryRow("select * from students where id=?", id)

	student := model.Student{}

	err := row.Scan(&student.ID, &student.Name, &student.Age)

	if err != nil {
		return student, err
	}

	return student, nil
}

func (dbInitializer *DBInitializer) SelectStudentsBasedName(name string) ([]model.Student, error) {

	rows, err := dbInitializer.db.Query("select * from students where name=?", name)

	if err != nil {
		log.Fatal(err.Error())
	}

	students := []model.Student{}

	for rows.Next() {

		student := model.Student{}

		err = rows.Scan(&student.ID, &student.Name, &student.Age)

		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		students = append(students, student)
	}

	return students, nil
}

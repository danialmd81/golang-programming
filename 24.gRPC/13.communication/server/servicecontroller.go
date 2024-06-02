package server

import (
	"communication/dbtools"
	"communication/model"
	"communication/protocol"
	"context"
)

type GrpcController struct {
	dbconnection *dbtools.DBInitializer
}

func GrpcServerInitializer(driverName string, dataSourceName string) (*GrpcController, error) {

	db, err := dbtools.Connect(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	return &GrpcController{
		dbconnection: db,
	}, err
}

func (server *GrpcController) GetStudentByID(context context.Context, id *protocol.SearchByID) (*protocol.Student, error) {

	student, err := server.dbconnection.SelectStudentBasedID(id.GetId())

	if err != nil {
		return nil, err
	}

	return convertToGrpcStudent(student), nil
}

func (server *GrpcController) GetStudentsByName(name *protocol.SearchByName, grpcStudents protocol.StudentService_GetStudentsByNameServer) error {

	students, err := server.dbconnection.SelectStudentsBasedName(name.GetName())

	if err != nil {
		return err
	}

	for _, student := range students {

		grpcStudent := convertToGrpcStudent(student)

		err := grpcStudents.Send(grpcStudent)

		if err != nil {
			return err
		}
	}
	return nil
}

func convertToGrpcStudent(student model.Student) *protocol.Student {

	return &protocol.Student{

		Id:   student.ID,
		Name: student.Name,
		Age:  student.Age,
	}
}

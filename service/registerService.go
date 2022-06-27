package service

import (
	"example-project/model"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DatabaseInterface
type DatabaseInterface interface {
	UpdateMany(docs []interface{}) interface{}
	UpdateOne(docs interface{}) interface{}
	GetByID(id string) model.Employee
	DeleteByID(id string) (*mongo.DeleteResult, error)
	GetAll() ([]model.Employee, error)
}

type EmployeeService struct {
	DbService DatabaseInterface
}

func NewEmployeeService(dbInterface DatabaseInterface) EmployeeService {
	return EmployeeService{
		DbService: dbInterface,
	}
}

func (s EmployeeService) CreateEmployees(employees []model.Employee) interface{} {

	var emp []interface{}
	for _, employee := range employees {
		emp = append(emp, employee)

	}
	return s.DbService.UpdateMany(emp)
}

func (s EmployeeService) GetEmployeeById(id string) model.Employee {
	return s.DbService.GetByID(id)
}

func (s EmployeeService) DeleteEmployeeById(id string) (*mongo.DeleteResult, error) {
	result, err := s.DbService.DeleteByID(id)
	return result, err

}

func (s EmployeeService) GetAllEmployees() ([]model.Employee, error) {
	result, err := s.DbService.GetAll()
	return result, err
}

func (s EmployeeService) CreateEmployee(employee model.Employee) interface{} {
	var emp interface{}
	emp = employee
	return s.DbService.UpdateOne(emp)
}

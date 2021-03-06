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
	GetPaginated(int, int) (model.PaginatedPayload, error)
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
	if len(employees) > 1 {
		for _, employee := range employees {
			emp = append(emp, employee)

		}
		return s.DbService.UpdateMany(emp)
	} else {
		return s.DbService.UpdateOne(employees[0])
	}

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

func (s EmployeeService) GetPaginatedEmployees(page int, limit int) (model.PaginatedPayload, error) {
	result, err := s.DbService.GetPaginated(page, limit)
	return result, err
}

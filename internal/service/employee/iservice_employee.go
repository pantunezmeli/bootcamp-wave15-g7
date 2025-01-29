package employee

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

type EmployeeService interface {
	// TODO
	FindAll() (employeesData map[int]dto.EmployeeDoc, err error)
	FindById(id int) (employeeData dto.EmployeeDoc, err error)
	New(employeeData dto.EmployeeDoc) (newEmployeeData dto.EmployeeDoc, err error)
	Edit(id int, employeeData dto.EmployeeDoc) (newEmployeeData dto.EmployeeDoc, err error)
	DeleteById(id int) (err error)
}

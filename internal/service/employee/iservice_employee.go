package employee

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

var ErrEmployeeNotFound = errors.New("employee not found")
var ErrWarehouseNotFound = errors.New("warehouse not found")
var ErrEmptyField = errors.New("employee data lacks a required field")
var ErrCardNumberAlreadyExists = errors.New("employee card number already exists")
var ErrInboundOrderNeedsEmployee = errors.New("inbound order linked to this employee")

type EmployeeService interface {
	// TODO
	FindAll() (employeesData []dto.EmployeeDoc, err error)
	FindById(id int) (employeeData dto.EmployeeDoc, err error)
	New(employeeData dto.EmployeeDoc) (newEmployeeData dto.EmployeeDoc, err error)
	Edit(id int, employeeData dto.EmployeeDoc) (newEmployeeData dto.EmployeeDoc, err error)
	DeleteById(id int) (err error)
}

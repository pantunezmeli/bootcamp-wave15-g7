package employee

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

type EmployeeService interface {
	// TODO
	FindAll() (employeesData map[int]dto.EmployeeDoc, err error)
	FindById() (err error)
	New() (err error)
	Update() (err error)
	DeleteById() (err error)
}

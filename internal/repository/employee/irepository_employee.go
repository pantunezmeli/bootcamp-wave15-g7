package employee

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"

type EmployeeRepository interface {
	// TODO
	FindAll() (employees map[int]model.Employee, err error)
	FindById(id int) (employee model.Employee, err error)
	New(employee model.Employee) (newEmployee model.Employee, err error)
	Update() (err error)
	DeleteById() (err error)
}

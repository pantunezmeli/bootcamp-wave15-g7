package employee

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"

type EmployeeRepository interface {
	// TODO
	FindAll() (employees map[int]model.Employee, err error)
	FindById() (err error)
	New() (err error)
	Update() (err error)
	DeleteById() (err error)
}

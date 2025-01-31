package employee

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type EmployeeRepository interface {
	// TODO
	FindAll() (employees map[int]models.Employee, err error)
	FindById(id int) (employee models.Employee, err error)
	New(employee models.Employee) (newEmployee models.Employee, err error)
	Edit(id int, employee models.Employee) (updatedEmployee models.Employee, err error)
	DeleteById(id int) (err error)
}

package employee

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

var ErrIdNotFound = errors.New("employee not found")
var ErrCardNumberNotUnique = errors.New("card number must be unique")

type EmployeeRepository interface {
	// TODO
	FindAll() (employees map[int]models.Employee, err error)
	FindById(id int) (employee models.Employee, err error)
	New(employee models.Employee) (newEmployee models.Employee, err error)
	Edit(id int, employee models.Employee) (updatedEmployee models.Employee, err error)
	DeleteById(id int) (err error)
}

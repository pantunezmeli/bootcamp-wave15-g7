package employee

import (
	"database/sql"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

func NewEmployeeSQL(database *sql.DB) *EmployeeSQL {
	return &EmployeeSQL{db: database}
}

type EmployeeSQL struct {
	db *sql.DB
}

func (r *EmployeeSQL) FindAll() (employees map[int]models.Employee, err error) {
	return
}

func (r *EmployeeSQL) FindById(id int) (employee models.Employee, err error) {
	return
}

func (r *EmployeeSQL) New(employee models.Employee) (newEmployee models.Employee, err error) {
	return
}

func (r *EmployeeSQL) Edit(id int, employee models.Employee) (updatedEmployee models.Employee, err error) {
	return
}

func (r *EmployeeSQL) DeleteById(id int) (err error) {
	return
}

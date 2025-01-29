package employee

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage"
)

var ErrIdNotFound = errors.New("employee not found")

func NewEmployeeMap(storage storage.EmployeeJSONFile) *EmployeeMap {
	return &EmployeeMap{st: storage}
}

type EmployeeMap struct {
	st storage.EmployeeJSONFile
}

func (r *EmployeeMap) FindAll() (employees map[int]model.Employee, err error) {
	file, err := r.st.Load()
	if err != nil {
		return
	}
	employees = make(map[int]model.Employee)

	// copy db
	for key, value := range file {
		employees[key] = value
	}

	return
}

func (r *EmployeeMap) FindById(id int) (employee model.Employee, err error) {
	file, err := r.st.Load()
	if err != nil {
		return
	}

	for _, value := range file {
		if value.Id.GetId() == id {
			employee = value
			return
		}
	}

	err = ErrIdNotFound
	return
}

func (r *EmployeeMap) New() (err error) {
	return nil
}

func (r *EmployeeMap) Update() (err error) {
	return nil
}

func (r *EmployeeMap) DeleteById() (err error) {
	return nil
}

package employee

import (
	"github.com/leofierens/bootcamp-wave15-g7/internal/storage"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
)

func NewEmployeeMap(storage storage.EmployeeJSONFile) *EmployeeMap {
	return &EmployeeMap{st: storage}
}

type EmployeeMap struct {
	st storage.EmployeeJSONFile
}

func (r *EmployeeMap) FindAll() (employees map[int]model.Employee, err error) {
	return nil, nil
}

func (r *EmployeeMap) FindById() (err error) {
	return nil
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

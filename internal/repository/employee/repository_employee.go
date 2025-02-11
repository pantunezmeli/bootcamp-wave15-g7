package employee

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	storage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/employee_storage"
)

func NewEmployeeMap(storage storage.EmployeeJSONFile) *EmployeeMap {
	return &EmployeeMap{st: storage}
}

type EmployeeMap struct {
	st storage.EmployeeJSONFile
}

func (r *EmployeeMap) FindAll() (employees map[int]models.Employee, err error) {
	file, err := r.st.Load()
	if err != nil {
		return
	}
	employees = make(map[int]models.Employee)

	for key, value := range file {
		employees[key] = value
	}

	return
}

func (r *EmployeeMap) FindById(id int) (employee models.Employee, err error) {
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

func (r *EmployeeMap) New(employee models.Employee) (newEmployee models.Employee, err error) {
	lastId, err := r.st.GetLastId()
	if err != nil {
		return
	}
	newEmployee = employee
	newEmployee.Id, err = value_objects.NewId(lastId + 1)
	if err != nil {
		return
	}
	err = r.st.Save(newEmployee)
	if err == storage.ErrCardNumberExists {
		err = ErrCardNumberNotUnique
	}
	return
}

func (r *EmployeeMap) Edit(id int, employee models.Employee) (updatedEmployee models.Employee, err error) {
	file, err := r.st.Load()
	if err != nil {
		return
	}

	for _, value := range file {
		if value.Id.GetId() == id {
			if err != nil {
				return
			}

			updatedEmployee = value

			if employee.CardNumber.GetCardNumber() != "" {
				updatedEmployee.CardNumber = employee.CardNumber
				if err = r.st.CheckCardNumber(employee.CardNumber.GetCardNumber()); err != nil {
					err = ErrCardNumberNotUnique
					return
				}
			}
			if employee.FirstName.GetName() != "" {
				updatedEmployee.FirstName = employee.FirstName
			}
			if employee.LastName.GetName() != "" {
				updatedEmployee.LastName = employee.LastName
			}
			if employee.WarehouseId.GetId() != 0 {
				updatedEmployee.WarehouseId = employee.WarehouseId
			}

			err = r.st.Erase(value)
			if err != nil {
				return
			}

			err = r.st.Save(updatedEmployee)
			return
		}
	}

	err = ErrIdNotFound
	return
}

func (r *EmployeeMap) DeleteById(id int) (err error) {
	file, err := r.st.Load()
	if err != nil {
		return
	}

	for _, value := range file {
		if value.Id.GetId() == id {
			err = r.st.Erase(value)
			return
		}
	}

	err = ErrIdNotFound
	return
}

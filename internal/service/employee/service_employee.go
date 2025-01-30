package employee

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/employee"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

var ErrEmployeeNotFound = errors.New("employee not found")
var ErrEmptyField = errors.New("employee data lacks a required field")

func NewDefaultService(repository employee.EmployeeRepository) *DefaultService {
	return &DefaultService{rp: repository}
}

type DefaultService struct {
	rp employee.EmployeeRepository
}

func (s *DefaultService) FindAll() (employeesData map[int]dto.EmployeeDoc, err error) {
	employeesFound, err := s.rp.FindAll()
	if err != nil {
		return
	}
	employeesData = make(map[int]dto.EmployeeDoc)
	for key, value := range employeesFound {
		employeesData[key] = dto.EmployeemodelsToDto(value)
	}
	return
}

func (s *DefaultService) FindById(id int) (employeeData dto.EmployeeDoc, err error) {
	employeeFound, errId := s.rp.FindById(id)
	if errId != nil {
		if errors.Is(errId, employee.ErrIdNotFound) {
			err = ErrEmployeeNotFound
		}
		return
	}
	employeeData = dto.EmployeemodelsToDto(employeeFound)
	return
}

func (s *DefaultService) New(employeeData dto.EmployeeDoc) (newEmployeeData dto.EmployeeDoc, err error) {
	if employeeData.CardNumber == "" || employeeData.FirstName == "" || employeeData.LastName == "" || employeeData.WarehouseId == 0 {
		err = ErrEmptyField
		return
	}

	employee, err := dto.EmployeeDtoTomodels(employeeData)
	if err != nil {
		return
	}

	newEmployee, err := s.rp.New(employee)
	if err != nil {
		return
	}

	newEmployeeData = dto.EmployeemodelsToDto(newEmployee)
	return
}

func (s *DefaultService) Edit(id int, employeeData dto.EmployeeDoc) (newEmployeeData dto.EmployeeDoc, err error) {
	_, errId := s.rp.FindById(id)
	if errId != nil {
		if errors.Is(errId, employee.ErrIdNotFound) {
			err = ErrEmployeeNotFound
		}
		return
	}

	employee, err := dto.EmployeeDtoToModeWithoutValidation(employeeData)
	if err != nil {
		return
	}
	updatedEmployee, err := s.rp.Edit(id, employee)
	if err != nil {
		return
	}

	newEmployeeData = dto.EmployeemodelsToDto(updatedEmployee)
	return
}

func (s *DefaultService) DeleteById(id int) (err error) {
	_, errId := s.rp.FindById(id)
	if errId != nil {
		if errors.Is(errId, employee.ErrIdNotFound) {
			err = ErrEmployeeNotFound
		}
		return
	}

	err = s.rp.DeleteById(id)
	if errors.Is(err, employee.ErrIdNotFound) {
		err = ErrEmployeeNotFound
	}
	return
}

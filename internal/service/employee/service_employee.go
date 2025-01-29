package employee

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/employee"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

var ErrEmployeeNotFound = errors.New("employee not found")

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
		employeesData[key] = dto.EmployeeModelToDto(value)
	}
	return
}

func (s *DefaultService) FindById(id int) (employeeData dto.EmployeeDoc, err error) {
	employeeFound, errId := s.rp.FindById(id)
	if errId != nil {
		if errors.Is(errId, employee.ErrIdNotFound) {
			err = ErrEmployeeNotFound
			return
		}
		return
	}
	employeeData = dto.EmployeeModelToDto(employeeFound)
	return
}

func (s *DefaultService) New() (err error) {
	return nil
}

func (s *DefaultService) Update() (err error) {
	return nil
}

func (s *DefaultService) DeleteById() (err error) {
	return nil
}

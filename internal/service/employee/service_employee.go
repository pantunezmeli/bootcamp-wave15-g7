package employee

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/employee"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

func NewDefaultService(repository employee.EmployeeRepository) *DefaultService {
	return &DefaultService{rp: repository}
}

type DefaultService struct {
	rp employee.EmployeeRepository
}

func (s *DefaultService) FindAll() (employeesData map[int]dto.EmployeeDoc, err error) {
	employeesFound, err := s.rp.FindAll()
	employeesData = make(map[int]dto.EmployeeDoc)
	for key, value := range employeesFound {
		employeesData[key] = dto.EmployeeModelToDto(value)
	}
	return
}

func (s *DefaultService) FindById() (err error) {
	return nil
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

package employee

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/employee"
)

func NewDefaultService(repository employee.EmployeeRepository) *DefaultService {
	return &DefaultService{rp: repository}
}

type DefaultService struct {
	rp employee.EmployeeRepository
}

func (s *DefaultService) FindAll() (employees map[int]model.Employee, err error) {
	employees, err = s.rp.FindAll()
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

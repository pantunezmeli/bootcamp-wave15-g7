package employee

import (
	"errors"
	"sort"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/employee"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/employee"
)

func NewDefaultService(repository employee.EmployeeRepository) *DefaultService {
	return &DefaultService{rp: repository}
}

type DefaultService struct {
	rp employee.EmployeeRepository
}

func (s *DefaultService) FindAll() (employeesData []dto.EmployeeDoc, err error) {
	employeesFound, err := s.rp.FindAll()
	if err != nil {
		return
	}
	employeesData = make([]dto.EmployeeDoc, 0, len(employeesFound))
	for _, value := range employeesFound {
		employeesData = append(employeesData, dto.EmployeeModelToDto(value))
	}
	sort.Slice(employeesData, func(i, j int) bool {
		return employeesData[i].Id < employeesData[j].Id
	})
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
	employeeData = dto.EmployeeModelToDto(employeeFound)
	return
}

func (s *DefaultService) New(employeeData dto.EmployeeDoc) (newEmployeeData dto.EmployeeDoc, err error) {
	if employeeData.CardNumber == "" || employeeData.FirstName == "" || employeeData.LastName == "" || employeeData.WarehouseId == 0 {
		err = ErrEmptyField
		return
	}

	employeeData.Id++
	employeeModel, err := dto.EmployeeDtoToModel(employeeData)
	if err != nil {
		return
	}

	newEmployee, err := s.rp.New(employeeModel)
	if err != nil {
		if err == employee.ErrCardNumberNotUnique {
			err = ErrCardNumberAlreadyExists
		}
		if err == employee.ErrWarehouseIdNotFound {
			err = ErrWarehouseNotFound
		}
		return
	}

	newEmployeeData = dto.EmployeeModelToDto(newEmployee)
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

	employeeModel, err := dto.EmployeeDtoToModelWithoutValidation(employeeData)
	if err != nil {
		return
	}
	updatedEmployee, err := s.rp.Edit(id, employeeModel)
	if err != nil {
		if errors.Is(err, employee.ErrIdNotFound) {
			err = ErrEmployeeNotFound
		}
		if err == employee.ErrCardNumberNotUnique {
			err = ErrCardNumberAlreadyExists
		}
		if err == employee.ErrWarehouseIdNotFound {
			err = ErrWarehouseNotFound
		}
		return
	}

	newEmployeeData = dto.EmployeeModelToDto(updatedEmployee)
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
	if errors.Is(err, employee.ErrInboundOrderFK) {
		err = ErrInboundOrderNeedsEmployee
	}
	if errors.Is(err, employee.ErrIdNotFound) {
		err = ErrEmployeeNotFound
	}
	return
}

func (s *DefaultService) ReportInboundOrders(id string) (reportsInboundOrders []dto.ReportInboundOrdersDoc, err error) {
	employees, inboundOrdersPerEmployee, err := s.rp.ReportInboundOrders(id)
	if errors.Is(err, employee.ErrIdNotFound) {
		err = ErrEmployeeNotFound
	}

	for _, value := range employees {
		var reportInboundOrders dto.ReportInboundOrdersDoc

		reportInboundOrders.Id = value.Id.GetId()
		reportInboundOrders.CardNumber = value.CardNumber.GetCardNumber()
		reportInboundOrders.FirstName = value.FirstName.GetName()
		reportInboundOrders.LastName = value.LastName.GetName()
		reportInboundOrders.WarehouseId = value.WarehouseId.GetId()

		reportInboundOrders.InboundOrders = inboundOrdersPerEmployee[reportInboundOrders.Id]

		reportsInboundOrders = append(reportsInboundOrders, reportInboundOrders)
	}
	return
}

package dto

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

type EmployeeDoc struct {
	Id          int    `json:"id"`
	CardNumber  string `json:"card_number_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	WarehouseId int    `json:"warehouse_id"`
}

func EmployeeDtoTomodels(dto EmployeeDoc) (employee models.Employee, err error) {
	newId, errValidation := domain.NewId(dto.Id)
	if errValidation != nil {
		return
	}
	newCardNumber, errValidation := domain.NewCardNumber(dto.CardNumber)
	if errValidation != nil {
		return
	}
	newFirstName, errValidation := domain.NewName(dto.FirstName)
	if errValidation != nil {
		return
	}
	newLastName, errValidation := domain.NewName(dto.LastName)
	if errValidation != nil {
		return
	}
	newWarehouseId, errValidation := domain.NewId(dto.WarehouseId)
	if errValidation != nil {
		return
	}
	employee = models.Employee{
		Id:          newId,
		CardNumber:  newCardNumber,
		FirstName:   newFirstName,
		LastName:    newLastName,
		WarehouseId: newWarehouseId,
	}
	return
}

func EmployeeDtoToModeWithoutValidation(dto EmployeeDoc) (employee models.Employee, err error) {
	employee = models.Employee{
		Id:          domain.NewOptionalId(dto.Id),
		CardNumber:  domain.NewOptionalCardNumber(dto.CardNumber),
		FirstName:   domain.NewOptionalName(dto.FirstName),
		LastName:    domain.NewOptionalName(dto.LastName),
		WarehouseId: domain.NewOptionalId(dto.WarehouseId),
	}
	return
}

func EmployeemodelsToDto(employee models.Employee) (dto EmployeeDoc) {
	dto = EmployeeDoc{
		Id:          employee.Id.GetId(),
		CardNumber:  employee.CardNumber.GetCardNumber(),
		FirstName:   employee.FirstName.GetName(),
		LastName:    employee.LastName.GetName(),
		WarehouseId: employee.WarehouseId.GetId(),
	}
	return
}

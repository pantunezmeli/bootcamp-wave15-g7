package dto

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
)

type EmployeeDoc struct {
	Id          int    `json:"id"`
	CardNumber  string `json:"card_number_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	WarehouseId int    `json:"warehouse_id"`
}

func EmployeeDtoToModel(dto EmployeeDoc) (employee model.Employee, err error) {
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
	employee = model.Employee{
		Id:          newId,
		CardNumber:  newCardNumber,
		FirstName:   newFirstName,
		LastName:    newLastName,
		WarehouseId: newWarehouseId,
	}
	return
}

func EmployeeModelToDto(employee model.Employee) (dto EmployeeDoc) {
	dto = EmployeeDoc{
		Id:          employee.Id.GetId(),
		CardNumber:  employee.CardNumber.GetCardNumber(),
		FirstName:   employee.FirstName.GetName(),
		LastName:    employee.LastName.GetName(),
		WarehouseId: employee.WarehouseId.GetId(),
	}
	return
}

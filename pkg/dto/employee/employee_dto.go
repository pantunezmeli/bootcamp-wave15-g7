package employee

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

type EmployeeDoc struct {
	Id          int    `json:"id"`
	CardNumber  string `json:"id_card_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	WarehouseId int    `json:"warehouse_id"`
}

type ReportInboundOrdersDoc struct {
	Id            int    `json:"id"`
	CardNumber    string `json:"id_card_number"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	WarehouseId   int    `json:"warehouse_id"`
	InboundOrders int    `json:"inbound_orders"`
}

func EmployeeDtoToModel(dto EmployeeDoc) (employee models.Employee, err error) {
	newId, errValidation := value_objects.NewId(dto.Id)
	if errValidation != nil {
		return
	}
	newCardNumber, errValidation := value_objects.NewCardNumber(dto.CardNumber)
	if errValidation != nil {
		return
	}
	newFirstName, errValidation := value_objects.NewName(dto.FirstName)
	if errValidation != nil {
		return
	}
	newLastName, errValidation := value_objects.NewName(dto.LastName)
	if errValidation != nil {
		return
	}
	newWarehouseId, errValidation := value_objects.NewId(dto.WarehouseId)
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

func EmployeeDtoToModelWithoutValidation(dto EmployeeDoc) (employee models.Employee, err error) {
	employee = models.Employee{
		Id:          value_objects.NewOptionalId(dto.Id),
		CardNumber:  value_objects.NewOptionalCardNumber(dto.CardNumber),
		FirstName:   value_objects.NewOptionalName(dto.FirstName),
		LastName:    value_objects.NewOptionalName(dto.LastName),
		WarehouseId: value_objects.NewOptionalId(dto.WarehouseId),
	}
	return
}

func EmployeeModelToDto(employee models.Employee) (dto EmployeeDoc) {
	dto = EmployeeDoc{
		Id:          employee.Id.GetId(),
		CardNumber:  employee.CardNumber.GetCardNumber(),
		FirstName:   employee.FirstName.GetName(),
		LastName:    employee.LastName.GetName(),
		WarehouseId: employee.WarehouseId.GetId(),
	}
	return
}

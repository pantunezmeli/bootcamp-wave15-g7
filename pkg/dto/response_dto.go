package dto

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

type SellerDoc struct {
	ID          *int    `json:"id"`
	Cid         *int    `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *string `json:"telephone"`
}

func ParseModelToDto(sellerModel models.Seller) (sellerDto SellerDoc) {
	sellerDto = SellerDoc{
		ID:          sellerModel.ID.Value(),
		Cid:         sellerModel.Cid.Value(),
		CompanyName: sellerModel.CompanyName.Value(),
		Address:     sellerModel.Address.Value(),
		Telephone:   sellerModel.Telephone.Value(),
	}
	return
}

func ParseDtoToModel(sellerDto SellerDoc) (sellerModel models.Seller, err error) {
	cid, err := value_objects.NewCid(*sellerDto.Cid)
	if err != nil {
		return
	}
	companyName, err := value_objects.NewCompanyName(*sellerDto.CompanyName)
	if err != nil {
		return
	}
	address, err := value_objects.NewSellerAddress(*sellerDto.Address)
	if err != nil {
		return
	}
	telephone, err := value_objects.NewSellerTelephone(*sellerDto.Telephone)
	if err != nil {
		return
	}

	sellerModel = models.Seller{
		SellerAttributes: models.SellerAttributes{
			Cid:         cid,
			CompanyName: companyName,
			Address:     address,
			Telephone:   telephone,
		},
	}
	return
}

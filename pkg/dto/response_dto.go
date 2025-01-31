package dto

import (
	"fmt"

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

func EmployeeDtoToModel(dto EmployeeDoc) (employee models.Employee, err error) {
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

func EmployeeDtoToModelWithoutValidation(dto EmployeeDoc) (employee models.Employee, err error) {
	employee = models.Employee{
		Id:          domain.NewOptionalId(dto.Id),
		CardNumber:  domain.NewOptionalCardNumber(dto.CardNumber),
		FirstName:   domain.NewOptionalName(dto.FirstName),
		LastName:    domain.NewOptionalName(dto.LastName),
		WarehouseId: domain.NewOptionalId(dto.WarehouseId),
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

type WareHouseDoc struct {
	Id                 int    `json:"id"`
	WareHouseCode      string `json:"warehouse_code"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}

func (w WareHouseDoc) ConvertToModel(req WareHouseDoc) (models.WareHouse, error) {

	wareHouseCode, err := domain.NewWareHouseCode(req.WareHouseCode)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse code '%s': %w", req.WareHouseCode, err)
	}

	address, err := domain.NewAddress(req.Address)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse address '%s': %w", req.Address, err)
	}

	telephone, err := domain.NewTelephone(req.Telephone)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse telephone '%s': %w", req.Telephone, err)
	}

	minCapacity, err := domain.NewMinimunCapacity(req.MinimunCapacity)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse minimun capacity '%d': %w", req.MinimunCapacity, err)
	}

	minTemp, err := domain.NewMinimunTemperature(req.MinimunTemperature)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse minimun temperature '%d': %w", req.MinimunTemperature, err)
	}

	wh := models.WareHouse{
		WareHouseCode:      wareHouseCode,
		Address:            address,
		Telephone:          telephone,
		MinimunCapacity:    minCapacity,
		MinimunTemperature: minTemp,
	}
	return wh, nil
}

func (w WareHouseDoc) ConvertToDTO(req models.WareHouse) (wh WareHouseDoc, err error) {
	return WareHouseDoc{
		Id:                 domain.Id.GetId(req.Id),                                                 // Convertir value object Id a int
		WareHouseCode:      domain.WareHouseCode.GetWareHouseCode(req.WareHouseCode),                // Convertir value object WareHouseCode a string
		Address:            domain.Address.GetAddress(req.Address),                                  // Convertir value object Address a string
		Telephone:          domain.Telephone.GetTelephone(req.Telephone),                            // Convertir value object Telephone a string
		MinimunCapacity:    domain.MinimunCapacity.GetMinimunCapacity(req.MinimunCapacity),          // Convertir value object MinimunCapacity a int
		MinimunTemperature: domain.MinimunTemperature.GetMinimunTemperature(req.MinimunTemperature), // Convertir value object MinimunTemperature a int
	}, nil
}

func (w WareHouseDoc) ConvertToModelPatch(req WareHouseDoc, existingWarehouse models.WareHouse) (models.WareHouse, error) {
	if req.WareHouseCode != "" {
		newCode, err := domain.NewWareHouseCode(req.WareHouseCode)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse code '%s': %w", req.WareHouseCode, err)
		}
		existingWarehouse.WareHouseCode = newCode
	}

	if req.Address != "" {
		newAddress, err := domain.NewAddress(req.Address)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse address '%s': %w", req.Address, err)
		}
		existingWarehouse.Address = newAddress
	}
	if req.Telephone != "" {
		newTelephone, err := domain.NewTelephone(req.Telephone)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse telephone '%s': %w", req.Telephone, err)
		}
		existingWarehouse.Telephone = newTelephone
	}
	if req.MinimunCapacity > 0 {
		newMinCapacity, err := domain.NewMinimunCapacity(req.MinimunCapacity)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse minimun capacity '%d': %w", req.MinimunCapacity, err)
		}
		existingWarehouse.MinimunCapacity = newMinCapacity
	}

	if req.MinimunTemperature > -100 {
		newMinTemp, err := domain.NewMinimunTemperature(req.MinimunTemperature)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse minimun temperature '%d': %w", req.MinimunTemperature, err)
		}
		existingWarehouse.MinimunTemperature = newMinTemp
	}
	return existingWarehouse, nil
}


type SellerDoc struct {
	ID  *int `json:"id"`
	Cid *int `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address *string `json:"address"`
	Telephone *string `json:"telephone"`
}


func ParseModelToDto(sellerModel models.Seller) (sellerDto SellerDoc){
	sellerDto = SellerDoc{
			ID: sellerModel.ID.Value(),
			Cid: sellerModel.Cid.Value(),
			CompanyName: sellerModel.CompanyName.Value(),
			Address: sellerModel.Address.Value(),
			Telephone: sellerModel.Telephone.Value(),
		}
	return
}

func ParseDtoToModel(sellerDto SellerDoc) (sellerModel models.Seller, err error){
	cid, err := domain.NewCid(*sellerDto.Cid)
	if err != nil{
		return
	}
	companyName, err := domain.NewCompanyName(*sellerDto.CompanyName)
	if err != nil {
		return
	}
	address, err := domain.NewSellerAddress(*sellerDto.Address)
	if err != nil{
		return
	}
	telephone, err := domain.NewSellerTelephone(*sellerDto.Telephone)
	if err != nil {
		return
	}

	sellerModel = models.Seller{
		SellerAttributes: models.SellerAttributes{
			Cid: cid,
			CompanyName: companyName,
			Address: address,
			Telephone: telephone,
		},
	}
	return
}


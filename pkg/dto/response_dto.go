package dto

import (
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

type EmployeeDoc struct {
	Id          int    `json:"id"`
	CardNumber  string `json:"card_number_id"`
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

type WareHouseDoc struct {
	Id                 int    `json:"id"`
	WareHouseCode      string `json:"warehouse_code"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}

func (w WareHouseDoc) ConvertToModel(req WareHouseDoc) (models.WareHouse, error) {

	wareHouseCode, err := value_objects.NewWareHouseCode(req.WareHouseCode)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse code '%s': %w", req.WareHouseCode, err)
	}

	address, err := value_objects.NewAddress(req.Address)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse address '%s': %w", req.Address, err)
	}

	telephone, err := value_objects.NewTelephone(req.Telephone)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse telephone '%s': %w", req.Telephone, err)
	}

	minCapacity, err := value_objects.NewMinimunCapacity(req.MinimunCapacity)
	if err != nil {
		return models.WareHouse{}, fmt.Errorf("invalid warehouse minimun capacity '%d': %w", req.MinimunCapacity, err)
	}

	minTemp, err := value_objects.NewMinimunTemperature(req.MinimunTemperature)
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
		Id:                 value_objects.Id.GetId(req.Id),                                                 // Convertir value object Id a int
		WareHouseCode:      value_objects.WareHouseCode.GetWareHouseCode(req.WareHouseCode),                // Convertir value object WareHouseCode a string
		Address:            value_objects.Address.GetAddress(req.Address),                                  // Convertir value object Address a string
		Telephone:          value_objects.Telephone.GetTelephone(req.Telephone),                            // Convertir value object Telephone a string
		MinimunCapacity:    value_objects.MinimunCapacity.GetMinimunCapacity(req.MinimunCapacity),          // Convertir value object MinimunCapacity a int
		MinimunTemperature: value_objects.MinimunTemperature.GetMinimunTemperature(req.MinimunTemperature), // Convertir value object MinimunTemperature a int
	}, nil
}

func (w WareHouseDoc) ConvertToModelPatch(req WareHouseDoc, existingWarehouse models.WareHouse) (models.WareHouse, error) {
	if req.WareHouseCode != "" {
		newCode, err := value_objects.NewWareHouseCode(req.WareHouseCode)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse code '%s': %w", req.WareHouseCode, err)
		}
		existingWarehouse.WareHouseCode = newCode
	}

	if req.Address != "" {
		newAddress, err := value_objects.NewAddress(req.Address)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse address '%s': %w", req.Address, err)
		}
		existingWarehouse.Address = newAddress
	}
	if req.Telephone != "" {
		newTelephone, err := value_objects.NewTelephone(req.Telephone)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse telephone '%s': %w", req.Telephone, err)
		}
		existingWarehouse.Telephone = newTelephone
	}
	if req.MinimunCapacity > 0 {
		newMinCapacity, err := value_objects.NewMinimunCapacity(req.MinimunCapacity)
		if err != nil {
			return models.WareHouse{}, fmt.Errorf("invalid warehouse minimun capacity '%d': %w", req.MinimunCapacity, err)
		}
		existingWarehouse.MinimunCapacity = newMinCapacity
	}

	if req.MinimunTemperature > -100 {
		newMinTemp, err := value_objects.NewMinimunTemperature(req.MinimunTemperature)
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
	cid, err := value_objects.NewCid(*sellerDto.Cid)
	if err != nil{
		return
	}
	companyName, err := value_objects.NewCompanyName(*sellerDto.CompanyName)
	if err != nil {
		return
	}
	address, err := value_objects.NewSellerAddress(*sellerDto.Address)
	if err != nil{
		return
	}
	telephone, err := value_objects.NewSellerTelephone(*sellerDto.Telephone)
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


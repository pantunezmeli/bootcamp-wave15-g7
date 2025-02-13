package carrier

import (
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	repository "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/carrier"
)

type CarrierDoc struct {
	Id          int    `json:"id"`
	Cid         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityId  int    `json:"locality_id"`
}

type CarrierByLocalityID struct {
	LocalityID    int    `json:"locality_id"`
	LocalityName  string `json:"locality_name"`
	CarriesAmount int    `json:"carries_count"`
}

func (w CarrierDoc) ConvertToModel(req CarrierDoc) (models.Carrier, error) {

	cid, err := value_objects.NewCarrierCid(req.Cid)
	if err != nil {
		return models.Carrier{}, err
	}

	companyName, err := value_objects.NewCarrierCompanyName(req.CompanyName)
	if err != nil {
		return models.Carrier{}, err
	}

	address, err := value_objects.NewAddress(req.Address)
	if err != nil {
		return models.Carrier{}, err
	}

	telephone, err := value_objects.NewTelephone(req.Telephone)
	if err != nil {
		return models.Carrier{}, err
	}

	localityID, err := value_objects.NewLocalityId(req.LocalityId)
	if err != nil {
		return models.Carrier{}, err
	}

	c := models.Carrier{
		Cid:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
		LocalityId:  localityID,
	}
	return c, nil
}

func (w CarrierDoc) ConvertToDTO(req models.Carrier) (c CarrierDoc, err error) {
	if req.Id.GetId() <= 0 {
		return CarrierDoc{}, fmt.Errorf("invalid Carrier ID: value is nil")
	}
	if req.Cid.GetCarrierCid() == "" {
		return CarrierDoc{}, fmt.Errorf("invalid CarrierCid: value is empty")
	}
	if req.CompanyName.GetCarrierCompanyName() == "" {
		return CarrierDoc{}, fmt.Errorf("invalid CompanyName: value is empty")
	}
	if req.Address.GetAddress() == "" {
		return CarrierDoc{}, fmt.Errorf("invalid Address: value is empty")
	}
	if req.Telephone.GetTelephone() == "" {
		return CarrierDoc{}, fmt.Errorf("invalid Telephone: value is empty")
	}
	if req.LocalityId.GetLocalityId() <= 0 {
		return CarrierDoc{}, fmt.Errorf("invalid LocalityId: value is nil")
	}
	return CarrierDoc{
		Id:          value_objects.Id.GetId(req.Id),
		Cid:         value_objects.CarrierCid.GetCarrierCid(req.Cid),
		CompanyName: value_objects.CarrierCompanyName.GetCarrierCompanyName(req.CompanyName),
		Address:     value_objects.Address.GetAddress(req.Address),
		Telephone:   value_objects.Telephone.GetTelephone(req.Telephone),
		LocalityId:  value_objects.LocalityId.GetLocalityId(req.LocalityId),
	}, nil
}

func (w CarrierByLocalityID) ConvertToDTO(req repository.CarrierByLocality) (c CarrierByLocalityID, err error) {
	if req.LocalityID <= 0 {
		return CarrierByLocalityID{}, fmt.Errorf("invalid Locality ID")
	}
	if len(req.LocalityName) <= 0 {
		return CarrierByLocalityID{}, fmt.Errorf("invalid Locality Name")
	}
	if req.CarriesAmount < 0 {
		return CarrierByLocalityID{}, fmt.Errorf("invalid carries amount")
	}
	return CarrierByLocalityID{
		LocalityID:    req.LocalityID,
		LocalityName:  req.LocalityName,
		CarriesAmount: req.CarriesAmount,
	}, nil
}

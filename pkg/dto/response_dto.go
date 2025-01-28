package dto

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

type SellerDoc struct {
	ID  int `json:"id"`
	Cid int `json:"cid"`
	CompanyName string `json:"company_name"`
	Address string `json:"address"`
	Telephone string `json:"telephone"`
	
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
	cid, err := domain.NewCid(sellerDto.Cid)
	if err != nil{
		return
	}
	companyName, err := domain.NewCompanyName(sellerDto.CompanyName)
	if err != nil {
		return
	}
	address, err := domain.NewAddress(sellerDto.Address)
	if err != nil{
		return
	}
	telephone, err := domain.NewTelephone(sellerDto.Telephone)
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
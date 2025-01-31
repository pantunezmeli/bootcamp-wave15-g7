package seller_dto

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

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


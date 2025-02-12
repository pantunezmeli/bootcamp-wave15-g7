package seller_dto

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

type SellerRequest struct {
	Cid *string `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address *string `json:"address"`
	Telephone *string `json:"telephone"`
}

type SellerDoc struct {
	ID  value_objects.SellerId `json:"id"`
	Cid value_objects.Cid `json:"cid"`
	CompanyName value_objects.CompanyName `json:"company_name"`
	Address value_objects.SellerAddress `json:"address"`
	Telephone value_objects.SellerTelephone `json:"telephone"`
}


func ParseModelToResponse(sellerModel models.Seller) (sellerDto SellerDoc){
	sellerDto = SellerDoc{
			ID: sellerModel.ID,
			Cid: sellerModel.Cid,
			CompanyName: sellerModel.CompanyName,
			Address: sellerModel.Address,
			Telephone: sellerModel.Telephone,
		}
	return
}

func ParseRequestToModel(sellerDto SellerRequest) (sellerModel models.Seller, err error){
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


package seller_dto

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	seller_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/seller"
	locality_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/locality"

)

type SellerRequest struct {
	Cid *string `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address *string `json:"address"`
	Telephone *string `json:"telephone"`
	LocalityId *int `json:"locality_id"`
}

type SellerDoc struct {
	ID  seller_vo.SellerId `json:"id"`
	Cid seller_vo.Cid `json:"cid"`
	CompanyName seller_vo.CompanyName `json:"company_name"`
	Address seller_vo.SellerAddress `json:"address"`
	Telephone seller_vo.SellerTelephone `json:"telephone"`
	LocalityId locality_vo.LocalityId `json:"locality_id"`

}


func ParseModelToResponse(sellerModel models.Seller) (sellerDto SellerDoc){
	sellerDto = SellerDoc{
			ID: sellerModel.ID,
			Cid: sellerModel.Cid,
			CompanyName: sellerModel.CompanyName,
			Address: sellerModel.Address,
			Telephone: sellerModel.Telephone,
			LocalityId: sellerModel.LocalityId,
		}
	return
}

func ParseRequestToModel(sellerDto SellerRequest) (sellerModel models.Seller, err error){
	cid, err := seller_vo.NewCid(*sellerDto.Cid)
	if err != nil{
		return
	}
	companyName, err := seller_vo.NewCompanyName(*sellerDto.CompanyName)
	if err != nil {
		return
	}
	address, err := seller_vo.NewSellerAddress(*sellerDto.Address)
	if err != nil{
		return
	}
	telephone, err := seller_vo.NewSellerTelephone(*sellerDto.Telephone)
	if err != nil {
		return
	}
	locality, err := locality_vo.NewLocalityId(*sellerDto.LocalityId)
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
		LocalityId: locality,
	}
	return
}


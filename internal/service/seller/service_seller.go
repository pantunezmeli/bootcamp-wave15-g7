package seller

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/seller"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

type SellerDefault struct {
	rp seller.SellerRepository
}

func NewSellerDefault(rp seller.SellerRepository) *SellerDefault {
	return &SellerDefault{rp}
}

func (s *SellerDefault) GetAll() (sellers []dto.SellerDoc, err error) {
	sellersModel, err := s.rp.GetAll()
	if err != nil {
		return
	}
	for _, sellerModel := range sellersModel{
		sellerDto := parseModelToDto(sellerModel)
		sellers = append(sellers, sellerDto)
	}
	return
}

func (s *SellerDefault) GetById(id int) (seller dto.SellerDoc, err error){
	sellerModel, err := s.rp.GetById(id)
	if err != nil {
		return
	}

	seller = parseModelToDto(sellerModel)
	return


}




func parseModelToDto(sellerModel models.Seller) (sellerDto dto.SellerDoc){
	sellerDto = dto.SellerDoc{
			ID: sellerModel.ID.Value(),
			Cid: sellerModel.Cid.Value(),
			CompanyName: sellerModel.CompanyName.Value(),
			Address: sellerModel.Address.Value(),
			Telephone: sellerModel.Telephone.Value(),
		}
	return
}
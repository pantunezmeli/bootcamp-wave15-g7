package seller

import (
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
		sellerDto := dto.ParseModelToDto(sellerModel)
		sellers = append(sellers, sellerDto)
	}
	return
}

func (s *SellerDefault) GetById(id int) (seller dto.SellerDoc, err error){
	sellerModel, err := s.rp.GetById(id)
	if err != nil {
		return
	}

	seller = dto.ParseModelToDto(sellerModel)
	return


}


func (s *SellerDefault) Save(reqBody dto.SellerDoc) (seller dto.SellerDoc, err error){
	model, err := dto.ParseDtoToModel(reqBody)
	if err != nil{
		return
	}

	resModel, err := s.rp.Save(model)
	if err != nil {
		return 
	}

	seller = dto.ParseModelToDto(resModel)

	return


}


func (s *SellerDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}





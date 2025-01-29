package seller

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
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
	if err = s.ValidateAllParameters(reqBody); err != nil {
		return
	}

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


func (s *SellerDefault) Update(reqBody dto.SellerDoc) (seller dto.SellerDoc, err error){
	sellerModel, err := s.rp.GetById(*reqBody.ID)
	if err != nil {
		return
	}

	if reqBody.Cid != nil {
		cid, err := domain.NewCid(*reqBody.Cid)
		if err != nil {
			return dto.SellerDoc{}, err
		}
		sellerModel.Cid = cid
	}

	if reqBody.CompanyName != nil {
		companyName, err := domain.NewCompanyName(*reqBody.CompanyName)
		if err != nil {
			return dto.SellerDoc{}, err
		}
		sellerModel.CompanyName = companyName
	}

	if reqBody.Address != nil {
		address, err := domain.NewAddress(*reqBody.Address)
		if err != nil {
			return dto.SellerDoc{}, err
		}
		sellerModel.Address = address
	}

	if reqBody.Telephone != nil {
		telephone, err := domain.NewTelephone(*reqBody.Telephone)
		if err != nil {
			return dto.SellerDoc{}, err
		}
		sellerModel.Telephone = telephone
	}

	sellerModel, err = s.rp.Update(sellerModel)

	if err != nil{
		return
	}

	seller = dto.ParseModelToDto(sellerModel)
	return


}


func (s *SellerDefault) ValidateAllParameters(reqBody dto.SellerDoc) (err error){
	if reqBody.Address == nil || reqBody.Cid == nil || reqBody.Telephone == nil ||reqBody.CompanyName == nil {
		err = ErrMissingParameters
		return
	}
	return
}

package seller

import (
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	seller_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/seller"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/seller"
	seller_dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/seller"
)

type SellerDefault struct {
	rp seller.SellerRepository
}

func NewSellerDefault(rp seller.SellerRepository) *SellerDefault {
	return &SellerDefault{rp}
}

func (s *SellerDefault) GetAll() (sellers []seller_dto.SellerDoc, err error) {
	sellersModel, err := s.rp.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, sellerModel := range sellersModel{
		sellerDto := seller_dto.ParseModelToResponse(sellerModel)
		sellers = append(sellers, sellerDto)
	}
	return
}

func (s *SellerDefault) GetById(id int) (seller seller_dto.SellerDoc, err error){
	sellerModel, err := s.rp.GetById(id)
	if err != nil {
		return
	}

	seller = seller_dto.ParseModelToResponse(sellerModel)
	return


}


func (s *SellerDefault) Save(reqBody seller_dto.SellerRequest) (seller seller_dto.SellerDoc, err error){
	if err = s.ValidateAllParameters(reqBody); err != nil {
		return
	}

	model, err := seller_dto.ParseRequestToModel(reqBody)
	if err != nil{
		err = &ErrInvalidParameter{err.Error()}
		return
	}

	resModel, err := s.rp.Save(model)
	if err != nil {
		return 
	}

	seller = seller_dto.ParseModelToResponse(resModel)

	return


}


func (s *SellerDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}


func (s *SellerDefault) Update(reqBody seller_dto.SellerRequest, id int) (seller seller_dto.SellerDoc, err error){
	sellerModel, err := s.rp.GetById(id)
	if err != nil {
		return
	}

	err = modifyAttributes(reqBody, &sellerModel)
	if err != nil{
		return
	}


	sellerModel, err = s.rp.Update(sellerModel)
	if err != nil{
		return
	}

	seller = seller_dto.ParseModelToResponse(sellerModel)
	return


}


func (s *SellerDefault) ValidateAllParameters(reqBody seller_dto.SellerRequest) (err error) {
	if reqBody.Address == nil {
		err = &ErrMissingParameters{AddressString}
		return
	}
	if reqBody.Cid == nil {
		err = &ErrMissingParameters{CidString}
		return
	}
	if reqBody.Telephone == nil {
		err = &ErrMissingParameters{TelephoneString}
		return
	}
	if reqBody.CompanyName == nil {
		err = &ErrMissingParameters{CompanyNameString}
		return
	}
	return
}

func modifyAttributes(reqBody seller_dto.SellerRequest, modelToModify *models.Seller) (err error) {
	if reqBody.Cid != nil {
		cid, err := seller_vo.NewCid(*reqBody.Cid)
		if err != nil {
			return &ErrInvalidParameter{err.Error()}
		}
		modelToModify.Cid = cid
	}

	if reqBody.CompanyName != nil {
		companyName, err := seller_vo.NewCompanyName(*reqBody.CompanyName)
		if err != nil {
			return &ErrInvalidParameter{err.Error()}
		}
		modelToModify.CompanyName = companyName
	}

	if reqBody.Address != nil {
		address, err := seller_vo.NewSellerAddress(*reqBody.Address)
		if err != nil {
			return &ErrInvalidParameter{err.Error()}
		}
		modelToModify.Address = address
	}

	if reqBody.Telephone != nil {
		telephone, err := seller_vo.NewSellerTelephone(*reqBody.Telephone)
		if err != nil {
			return &ErrInvalidParameter{err.Error()}
		}
		modelToModify.Telephone = telephone
	}
	return
}
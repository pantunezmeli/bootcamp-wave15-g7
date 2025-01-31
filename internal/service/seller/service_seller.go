package seller

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
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
		return
	}
	for _, sellerModel := range sellersModel{
		sellerDto := seller_dto.ParseModelToDto(sellerModel)
		sellers = append(sellers, sellerDto)
	}
	return
}

func (s *SellerDefault) GetById(id int) (seller seller_dto.SellerDoc, err error){
	sellerModel, err := s.rp.GetById(id)
	if err != nil {
		return
	}

	seller = seller_dto.ParseModelToDto(sellerModel)
	return


}


func (s *SellerDefault) Save(reqBody seller_dto.SellerDoc) (seller seller_dto.SellerDoc, err error){
	if err = s.ValidateAllParameters(reqBody); err != nil {
		return
	}

	model, err := seller_dto.ParseDtoToModel(reqBody)
	if err != nil{
		err = &ErrInvalidParameter{err.Error()}
		return
	}

	resModel, err := s.rp.Save(model)
	if err != nil {
		return 
	}

	seller = seller_dto.ParseModelToDto(resModel)

	return


}


func (s *SellerDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}


func (s *SellerDefault) Update(reqBody seller_dto.SellerDoc) (seller seller_dto.SellerDoc, err error){
	sellerModel, err := s.rp.GetById(*reqBody.ID)
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

	seller = seller_dto.ParseModelToDto(sellerModel)
	return


}


func (s *SellerDefault) ValidateAllParameters(reqBody seller_dto.SellerDoc) (err error) {
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

func modifyAttributes(reqBody seller_dto.SellerDoc, modelToModify *models.Seller) (err error) {
	if reqBody.Cid != nil {
		cid, err := value_objects.NewCid(*reqBody.Cid)
		if err != nil {
			return &ErrInvalidParameter{err.Error()}
		}
		modelToModify.Cid = cid
	}

	if reqBody.CompanyName != nil {
		companyName, err := value_objects.NewCompanyName(*reqBody.CompanyName)
		if err != nil {
			return &ErrInvalidParameter{err.Error()}
		}
		modelToModify.CompanyName = companyName
	}

	if reqBody.Address != nil {
		address, err := value_objects.NewSellerAddress(*reqBody.Address)
		if err != nil {
			return &ErrInvalidParameter{err.Error()}
		}
		modelToModify.Address = address
	}

	if reqBody.Telephone != nil {
		telephone, err := value_objects.NewSellerTelephone(*reqBody.Telephone)
		if err != nil {
			return &ErrInvalidParameter{err.Error()}
		}
		modelToModify.Telephone = telephone
	}
	return
}
package locality

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/locality"
	locality_dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/locality"
	
	
)

type LocalityDefault struct {
	rp locality.LocalityRepository
}

func NewLocalityDefault(rp locality.LocalityRepository) *LocalityDefault{
	return &LocalityDefault{rp}
}

func (s *LocalityDefault) Save(reqBody locality_dto.LocalityRequest) (localityDto locality_dto.LocalityDoc, err error) {
	if err = s.ValidateAllParameters(reqBody); err != nil {
		return
	}

	model, err := locality_dto.ParseRequestToModel(reqBody)
	if err != nil{
		err = &ErrInvalidParameter{err.Error()}
		return
	}

	resModel, err := s.rp.Save(model)
	if err != nil {
		return 
	}

	localityDto = locality_dto.ParseModelToResponse(resModel)

	return

}

func (s *LocalityDefault) GetById(id int) (localityDto locality_dto.LocalityDoc, err error){
	localityModel, err := s.rp.GetById(id)
	if err != nil {
		return
	}

	localityDto = locality_dto.ParseModelToResponse(localityModel)
	return
}


func (s *LocalityDefault) GetReportSellers(id *int) (reports []locality_dto.SellerReport, err error){
	reports, err = s.rp.GetReportSellers(id)
	return

}

func (s *LocalityDefault) ValidateAllParameters(reqBody locality_dto.LocalityRequest) (err error) {
	if reqBody.Name == nil {
		err = &ErrMissingParameters{NameString}
		return
	}

	if reqBody.ProvinceId == nil {
		err = &ErrMissingParameters{ProvinceIdString}
		return
	}
	return
}


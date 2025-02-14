package inboundorder

import (
	repo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/inboundorder"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/inboundorder"
)

func NewDefaultService(repository repo.InboundOrderRepository) *DefaultService {
	return &DefaultService{rp: repository}
}

type DefaultService struct {
	rp repo.InboundOrderRepository
}

func (s *DefaultService) New(inboundOrderData dto.InboundOrderDoc) (newInboundOrderData dto.InboundOrderDoc, err error) {
	if err = validateFields(inboundOrderData); err != nil {
		return
	}

	inboundOrderData.Id++
	inboundOrderModel, err := dto.InboundOrderDtoToModel(inboundOrderData)
	if err != nil {
		return
	}

	newInboundOrder, err := s.rp.New(inboundOrderModel)
	if err != nil {
		if err == repo.ErrOrderNumberNotUnique {
			err = ErrOrderNumberAlreadyExists
		}
		if err == repo.ErrEmployeeIdNotFound {
			err = ErrEmployeeNotFound
		}
		if err == repo.ErrProductBatchIdNotFound {
			err = ErrProductBatchNotFound
		}
		if err == repo.ErrWarehouseIdNotFound {
			err = ErrWarehouseNotFound
		}
		return
	}

	newInboundOrderData = dto.InboundOrderModelToDto(newInboundOrder)
	return
}

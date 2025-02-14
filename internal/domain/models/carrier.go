package models

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

type Carrier struct {
	Id          value_objects.Id
	Cid         value_objects.CarrierCid
	CompanyName value_objects.CarrierCompanyName
	Address     value_objects.Address
	Telephone   value_objects.Telephone
	LocalityId  value_objects.LocalityId
}

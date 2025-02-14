package value_objects

import (
	"errors"
)

var (
	ErrInvalidLocalityId  = errors.New("invalid locality id")
	ErrInvalidCompanyName = errors.New("invalid company name")
)

// ############################# Locality ID ######################################
type LocalityId struct {
	value int
}

func NewLocalityId(value int) (localityId LocalityId, err error) {
	if value <= 0 {
		return LocalityId{}, ErrInvalidLocalityId
	}
	return LocalityId{value: value}, nil
}

func (localityId LocalityId) GetLocalityId() int {
	return localityId.value
}

// ############################# Carrier Cid ######################################
type CarrierCid struct {
	value string
}

func NewCarrierCid(value string) (cid CarrierCid, err error) {
	if len(value) <= 0 {
		return CarrierCid{}, ErrInvalidCid
	}
	return CarrierCid{value: value}, nil
}

func (cid CarrierCid) GetCarrierCid() string {
	return cid.value
}

// ############################# Company Name ######################################
type CarrierCompanyName struct {
	value string
}

func NewCarrierCompanyName(value string) (companyName CarrierCompanyName, err error) {
	if len(value) <= 0 {
		return CarrierCompanyName{}, ErrInvalidCompanyName
	}
	return CarrierCompanyName{value: value}, nil
}

func (companyName CarrierCompanyName) GetCarrierCompanyName() string {
	return companyName.value
}

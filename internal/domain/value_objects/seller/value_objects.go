package seller

import (
	"errors"
	"regexp"
	"strings"
)

var (
	CompanyNameMin      = 2
	CompanyNameMax      = 100
	AddressMin          = 2
	AddressMax          = 255
	TelephoneMin        = 2
	TelephoneMax        = 30
	validTelephoneRegex = `^[0-9+\-()\s]+$`

	ErrSellerInvalidId        = errors.New("id should be a positive number")
	ErrInvalidCid             = errors.New("cid should be a positive number")
	ErrCompanyNameTooShort    = errors.New("company name must be at least 2 characters long")
	ErrCompanyNameTooLong     = errors.New("company name must not exceed 100 characters")
	ErrAddressTooShort        = errors.New("address must be at least 2 characters long")
	ErrAddressTooLong         = errors.New("address must not exceed 255 characters")
	ErrTelephoneTooShort      = errors.New("phone number must be at least 8 characters long")
	ErrTelephoneTooLong       = errors.New("phone number must not exceed 30 characters")
	ErrSellerInvalidTelephone = errors.New("telephone contains invalid characters")
)

type SellerId int
func NewSellerId(value int) (id SellerId, err error) {
	if value <= 0 {
		err = ErrSellerInvalidId
		return
	}
	id = SellerId(value)
	return
}


type Cid string

func NewCid(value string) (cid Cid, err error) {
	if len(value) <= 1 {
		err = ErrInvalidCid
	}
	cid = Cid(value)
	return
}


type CompanyName string

func NewCompanyName(value string) (companyName CompanyName, err error) {
	value = strings.TrimSpace(value)
	if err = validateLength(value, CompanyNameMin, CompanyNameMax, ErrCompanyNameTooShort, ErrCompanyNameTooLong); err != nil {
		return
	}

	companyName = CompanyName(value)
	return
}

type SellerAddress string

func NewSellerAddress(value string) (address SellerAddress, err error) {
	if err = validateLength(value, AddressMin, AddressMax, ErrAddressTooShort, ErrAddressTooLong); err != nil {
		return
	}
	address = SellerAddress(value)
	return

}

type SellerTelephone string

func NewSellerTelephone(value string) (telephone SellerTelephone, err error) {
	if err = validateLength(value, TelephoneMin, TelephoneMax, ErrTelephoneTooShort, ErrTelephoneTooLong); err != nil {
		return
	}

	if err = validateTelephone(value); err != nil {
		return
	}
	telephone = SellerTelephone(value)
	return
}

func validateLength(value string, min, max int, errMin, errMax error) error {
	if len(value) < min {
		return errMin
	}
	if len(value) > max {
		return errMax
	}
	return nil
}

func validateTelephone(value string) error {
	validTelephone := regexp.MustCompile(validTelephoneRegex)
	if !validTelephone.MatchString(value) {
		return ErrSellerInvalidTelephone
	}
	return nil
}


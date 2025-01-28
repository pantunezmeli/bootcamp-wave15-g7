package domain

import (
	"errors"
	"regexp"
	"strings"
)

//los errores estan hardcodeados
var (
	CompanyNameMin            = 2
	CompanyNameMax            = 100
	AddressMin                = 2
	AddressMax                = 255
	TelephoneMin              = 2
	TelephoneMax              = 30
	validTelephoneRegex       = `^[0-9+\-()\s]+$`
        

	ErrInvalidId              = errors.New("id should be a positive number")
	ErrInvalidCid             = errors.New("cid should be a positive number")
	ErrCompanyNameTooShort    = errors.New("company name must be at least 2 characters long")
	ErrCompanyNameTooLong     = errors.New("company name must not exceed 100 characters")
	ErrAddressTooShort        = errors.New("address must be at least 2 characters long")
	ErrAddressTooLong         = errors.New("address must not exceed 255 characters")
	ErrTelephoneTooShort      = errors.New("phone number must be at least 8 characters long")
	ErrTelephoneTooLong       = errors.New("phone number must not exceed 30 characters")
	ErrInvalidTelephone       = errors.New("telephone contains invalid characters")

)


type Id struct {
	value int
}

func NewId(value int) (id Id, err error) {
	if value <= 0 {
		err = ErrInvalidId
		return
	}
	id = Id{value: value}
	return
}

func (id Id) Value() int {
	return id.value
}

type Cid struct {
	value int
}

func NewCid(value int) (cid Cid, err error){
	if value <= 0 {
		err = ErrInvalidCid
		return
	}
	cid = Cid{value: value}
	return
}

func (c Cid) Value() int {
	return c.value
}


type CompanyName struct {
	value string
}

func NewCompanyName(value string) (companyName CompanyName, err error){
	value = strings.TrimSpace(value)
	if err = validateLength(value, CompanyNameMin, CompanyNameMax, ErrCompanyNameTooShort, ErrCompanyNameTooLong); err != nil {
        return
    }

	//validar caracteres?

	companyName = CompanyName{value}
	return
}

func (c CompanyName) Value() string{
	return c.value
}

type Address struct {
	value string
}

func NewAddress(value string) (address Address, err error){
	if err = validateLength(value, AddressMin, AddressMax, ErrAddressTooShort, ErrAddressTooLong); err != nil {
        return
    }
	address = Address{value}
	return
	
}

func (a Address) Value() string {
	return a.value
}


type Telephone struct {
	value string
}

func NewTelephone(value string) (telephone Telephone, err error){
	if err = validateLength(value, TelephoneMin, TelephoneMax, ErrTelephoneTooShort, ErrTelephoneTooLong); err != nil {
        return
    }

	if err = validateTelephone(value); err != nil{
		return
	}
	telephone = Telephone{value}
	return
}


func (t Telephone) Value() string {
	return t.value
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
        return ErrInvalidTelephone
    }
    return nil
}
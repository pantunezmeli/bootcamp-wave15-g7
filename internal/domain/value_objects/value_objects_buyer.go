package value_objects

import (
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type CardNumberId struct {
	value int
}

type FirstName struct {
	first_Name string
}

type LastName struct {
	last_Name string
}

func NewCardNumberId(id int) (cardNumber CardNumberId, err error) {
	if id <= 0 {
		return CardNumberId{}, errorbase.ErrInvalidId
	}
	return CardNumberId{value: id}, nil
}

func (c CardNumberId) GetCardNumberId() int {
	return c.value
}

func NewFirstName(firstName string) (FirstName, error) {
	if firstName == "" {
		return FirstName{}, errorbase.ErrEmptyParameters
	}
	return FirstName{first_Name: firstName}, nil
}

func (f FirstName) GetFirstName() string {
	return f.first_Name
}

func NewLastName(lastName string) (LastName, error) {
	if lastName == "" {
		return LastName{}, errorbase.ErrEmptyParameters
	}
	return LastName{last_Name: lastName}, nil
}

func (l LastName) GetLastName() string {
	return l.last_Name
}

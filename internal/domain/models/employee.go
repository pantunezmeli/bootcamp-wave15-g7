package models

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"

type Employee struct {
	Id          domain.Id
	CardNumber  domain.CardNumber
	FirstName   domain.Name
	LastName    domain.Name
	WarehouseId domain.Id
}

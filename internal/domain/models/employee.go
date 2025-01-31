package models

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"

type Employee struct {
	Id          value_objects.Id
	CardNumber  value_objects.CardNumber
	FirstName   value_objects.Name
	LastName    value_objects.Name
	WarehouseId value_objects.Id
}

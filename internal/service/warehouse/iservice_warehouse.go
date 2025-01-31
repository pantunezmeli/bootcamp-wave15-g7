package warehouse

import (
	"fmt"

	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/warehouse"
)

type ErrInvalidParameter struct {
	Parameter string
}

func (e ErrInvalidParameter) Error() string {
	return fmt.Sprintf("missing parameter: %s", e.Parameter)
}

type IWareHouseService interface {
	// ! 1)
	FindAll() (w map[int]dto.WareHouseDoc, err error)

	// ! 2)
	GetWareHouseById(id int) (w dto.WareHouseDoc, err error)

	// ! 3)
	AddWareHouse(req dto.WareHouseDoc) (w dto.WareHouseDoc, err error)

	// ! 4)
	EditWareHouse(id int, req dto.WareHouseDoc) (wh dto.WareHouseDoc, err error)

	// ! 5)
	DeleteWarehouse(id int) error
}

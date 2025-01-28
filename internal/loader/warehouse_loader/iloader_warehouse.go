package warehouse

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type IWareHouseLoader interface {
	Load() (w map[int]models.WareHouse, err error)
}

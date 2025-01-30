package warehouse_storage

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

type IWareHouseStorage interface {
	Load() (w map[int]models.WareHouse, err error)
	Save(wh map[int]models.WareHouse) error
}

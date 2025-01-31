package seller_storage

import ("github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
"errors"
)

var (
	ErrSavingFile = errors.New("error saving the file")
	ErrParsingData = errors.New("error parsing data to bytes")
)



type SellerStorage interface {
	Load() (v map[int]models.Seller, err error)
	Save(map[int]models.Seller) (err error)
}

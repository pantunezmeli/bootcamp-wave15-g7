package loader

import ("github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
"errors"
)

var (
	ErrSavingFile = errors.New("error saving the file")
	ErrParsingData = errors.New("error parsing data to bytes")
)



type SellerLoader interface {
	Load() (v map[int]models.Seller, err error)
}

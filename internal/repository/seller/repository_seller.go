package seller

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/loader"

type SellerStorage struct {
	loader loader.SellerJSONFile 
}

func NewSellerStorage(loader loader.SellerJSONFile) *SellerStorage {
	return &SellerStorage{loader}
}


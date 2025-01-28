package seller

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/seller"

type SellerDefault struct {
	rp seller.SellerRepository
}

func NewSellerDefault(rp seller.SellerRepository) *SellerDefault {
	return &SellerDefault{rp}
}
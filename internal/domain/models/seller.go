package models

import ( seller_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/seller"
		locality_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/locality"
)
type SellerAttributes struct {
	Cid seller_vo.Cid

	CompanyName seller_vo.CompanyName

	Address seller_vo.SellerAddress

	Telephone seller_vo.SellerTelephone

}

type Seller struct {
	ID seller_vo.SellerId

	SellerAttributes

	LocalityId locality_vo.LocalityId

}
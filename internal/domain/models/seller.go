package models

import (
	v "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
)

type SellerAttributes struct {
	Cid v.Cid

	CompanyName v.CompanyName

	Address v.Address

	Telephone v.Telephone

}

type Seller struct {
	ID v.Id

	SellerAttributes
}
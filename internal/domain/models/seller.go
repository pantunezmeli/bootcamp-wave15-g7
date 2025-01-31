package models

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"

type SellerAttributes struct {
	Cid value_objects.Cid

	CompanyName value_objects.CompanyName

	Address value_objects.SellerAddress

	Telephone value_objects.SellerTelephone

}

type Seller struct {
	ID value_objects.SellerId

	SellerAttributes
}
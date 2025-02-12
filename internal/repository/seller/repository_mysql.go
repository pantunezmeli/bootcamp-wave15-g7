package seller

import (
	"database/sql"
	"fmt"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

type SellerMySql struct {
	db *sql.DB
}

func NewSellerMySql(db *sql.DB) *SellerMySql {
	return &SellerMySql{db}
}

func (s *SellerMySql) GetAll() (sellers []models.Seller, err error) {
	rows, err := s.db.Query(`
	SELECT id, cid, company_name, address, telephone
	from sellers
	`)
	if err != nil {
		err = ErrConnectionError
		return
	}

	for rows.Next(){
		var seller models.Seller
		if err = rows.Scan(&seller.ID, &seller.Cid, &seller.CompanyName, &seller.Address, &seller.Telephone); err != nil {
			fmt.Println(err)
			err = ErrConnectionError
			return
		}

		// idModel, err := value_objects.NewSellerId(id)
		// if err != nil {
		// 	return []models.Seller{}, err
		// }
		// cidModel, err := value_objects.NewCid(cid)
		// if err != nil {
		// 	return []models.Seller{}, err
		// }
		// companyNameModel, err := value_objects.NewCompanyName(companyName)
		// if err != nil {
		// 	return []models.Seller{}, err
		// }
		// addressModel, err := value_objects.NewSellerAddress(address)
		// if err != nil {
		// 	return []models.Seller{}, err
		// }
		// telephoneModel, err := value_objects.NewSellerTelephone(telephone)
		// if err != nil {
		// 	return []models.Seller{}, err
		// }

		sellers = append(sellers, seller)
	}
	if err = rows.Err(); err != nil{
		err = ErrConnectionError
	}
	return
}

func  (s *SellerMySql) GetById(id int) (seller models.Seller, err error) {
	return
}

func (s *SellerMySql) Save(modelWithoutId models.Seller) (seller models.Seller, err error) {
	return
}

func (s *SellerMySql) Delete(id int) (err error){
	return



}

func(s *SellerMySql) Update(sellerModel models.Seller) (sellerUpdated models.Seller, err error){
	return
}


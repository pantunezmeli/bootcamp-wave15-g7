package seller

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	seller_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/seller"

)

type SellerMySql struct {
	db *sql.DB
}

func NewSellerMySql(db *sql.DB) *SellerMySql {
	return &SellerMySql{db}
}

func (s *SellerMySql) GetAll() (sellers []models.Seller, err error) {
	rows, err := s.db.Query(`
	SELECT id, cid, company_name, address, telephone, locality_id
	from sellers
	`)
	if err != nil {
		err = ErrConnection
		return
	}
	defer rows.Close()


	for rows.Next(){
		var seller models.Seller
		if err = rows.Scan(&seller.ID, &seller.Cid, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityId); err != nil {
			fmt.Println(err)
			err = ErrConnection
			return
		}

		sellers = append(sellers, seller)
	}
	if err = rows.Err(); err != nil{
		err = ErrConnection
	}
	return
}

func  (s *SellerMySql) GetById(id int) (seller models.Seller, err error) {
	row := s.db.QueryRow(
		`
		SELECT id, cid, company_name, address, telephone, locality_id
		from sellers
		where id = ?
		`, id)

	if err = row.Scan(&seller.ID, &seller.Cid, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityId); err != nil{
		if errors.Is(err, sql.ErrNoRows){
			err = ErrSellerNotFound
			return
		}
		err = ErrConnection
	}

	return
}

func (s *SellerMySql) Save(modelWithoutId models.Seller) (seller models.Seller, err error) {
	result, err := s.db.Exec(`
	INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES
	(?,?,?,?,?)`, modelWithoutId.Cid, modelWithoutId.CompanyName, modelWithoutId.Address, modelWithoutId.Telephone, modelWithoutId.LocalityId)
	if err != nil {
		err = mapMySQLError(err)
		return
	}

	seller = modelWithoutId
	id, err := result.LastInsertId()
	if err != nil {
		err = ErrConnection
		return
	}
	seller.ID, err = seller_vo.NewSellerId(int(id))
	if err != nil {
		return
	}
	return
}

func (s *SellerMySql) Delete(id int) (err error){
	result, err := s.db.Exec(`DELETE FROM sellers WHERE id = ?`, id)
	if err != nil {
		err = ErrConnection
		return
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		err = ErrConnection
		return
	}
	if affectedRows == 0 {
		err = ErrSellerNotFound
		return
	}
	return
}

func(s *SellerMySql) Update(sellerModel models.Seller) (sellerUpdated models.Seller, err error){
	var exists bool
	err = s.db.QueryRow("SELECT 1 FROM sellers WHERE id = ?", sellerModel.ID).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrSellerNotFound
		} else {
			err = ErrConnection
		}
		return
	}

	_, err = s.db.Exec(`
	UPDATE sellers SET cid = ?, company_name = ?, address = ?, telephone = ?, locality_id = ? WHERE id = ?`,
	sellerModel.Cid, sellerModel.CompanyName, sellerModel.Address, sellerModel.Telephone, sellerModel.LocalityId, sellerModel.ID)
	if err != nil {
		err = mapMySQLError(err)
		return
	}

	sellerUpdated = sellerModel
	return
}

func mapMySQLError(err error) error {
    var sqlError *mysql.MySQLError
    if errors.As(err, &sqlError) {
        switch sqlError.Number {
        case 1452:
            return ErrLocalityNotFound
        case 1062:
            return ErrCidAlreadyExists
        }
    }
    return ErrConnection
}

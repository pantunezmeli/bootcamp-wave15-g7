package buyer

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	//"dario.cat/mergo"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"

	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/querys"
)

type BuyerRepository struct {
	db *sql.DB
}

func NewBuyerRepository(db *sql.DB) *BuyerRepository {
	return &BuyerRepository{db: db}
}

func (buyer *BuyerRepository) GetAll() ([]models.Buyer, error) {

	rows, err := buyer.db.Query(querys.SelectAllBuyers)
	if err != nil {
		return nil, errorbase.ErrDatabaseOperationFailed
	}

	defer rows.Close()

	var buyers []models.Buyer
	for rows.Next() {
		var buyer models.Buyer
		err := rows.Scan(&buyer.Id, &buyer.Card_Number_Id, &buyer.First_Name, &buyer.Last_Name)
		if err != nil {
			return nil, errorbase.ErrDatabaseOperationFailed
		}
		buyers = append(buyers, buyer)
	}

	if err = rows.Err(); err != nil {
		return nil, errorbase.ErrDatabaseOperationFailed
	}

	return buyers, nil
}

func (buyer *BuyerRepository) GetById(id int) (models.Buyer, error) {
	row := buyer.db.QueryRow(querys.SelectByID, id)
	var buyer_model models.Buyer

	err := row.Scan(&buyer_model.Id, &buyer_model.Card_Number_Id, &buyer_model.First_Name, &buyer_model.Last_Name)

	if err != nil {
		return models.Buyer{}, errorbase.ErrDatabaseOperationFailed
	}

	return buyer_model, nil
}

func (buyer *BuyerRepository) Create(entity models.Buyer) (models.Buyer, error) {

	if exist := buyer.cardIdExist(0, entity.Card_Number_Id); exist {
		return models.Buyer{}, errorbase.ErrConflict
	}

	attributes, err := buyer.validatemodels(entity)

	if err != nil {
		return models.Buyer{}, errorbase.ErrEmptyParameters
	}

	entity.Card_Number_Id = attributes["CardNumber"].(value_objects.CardNumberId).GetCardNumberId()
	entity.First_Name = attributes["FirstName"].(value_objects.FirstName).GetFirstName()
	entity.Last_Name = attributes["LastName"].(value_objects.LastName).GetLastName()

	result, err := buyer.db.Exec(querys.InsertBuyer, entity.Card_Number_Id, entity.First_Name, entity.Last_Name)
	if err != nil {
		return models.Buyer{}, errorbase.ErrDatabaseOperationFailed
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Buyer{}, errorbase.ErrDatabaseOperationFailed
	}

	entity.Id = int(id)

	return entity, nil
}

func (buyer *BuyerRepository) Update(id int, entity models.Buyer) (models.Buyer, error) {

	cardId := strings.TrimSpace(entity.Card_Number_Id)
	if cardId != "" {
		if exist := buyer.cardIdExist(id, entity.Card_Number_Id); exist {
			return models.Buyer{}, errorbase.ErrConflict
		}
	}

	cardNumberId, _ := value_objects.NewCardNumberId(entity.Card_Number_Id)
	firstName, _ := value_objects.NewFirstName(entity.First_Name)
	lastName, _ := value_objects.NewLastName(entity.Last_Name)

	var dinamicQuery []string
	var values []interface{}

	if cardNumberId.GetCardNumberId() != "" {
		dinamicQuery = append(dinamicQuery, "id_card_number = ?")
		values = append(values, cardNumberId.GetCardNumberId())
	}

	if firstName.GetFirstName() != "" {
		dinamicQuery = append(dinamicQuery, "first_name = ?")
		values = append(values, firstName.GetFirstName())
	}

	if lastName.GetLastName() != "" {
		dinamicQuery = append(dinamicQuery, "last_name = ?")
		values = append(values, lastName.GetLastName())
	}

	if len(dinamicQuery) == 0 {
		return models.Buyer{}, errorbase.ErrUnprocessable
	}

	query := fmt.Sprintf("UPDATE buyers SET %s WHERE id = ?", strings.Join(dinamicQuery, ", "))
	values = append(values, id)

	_, err := buyer.db.Exec(query, values...)
	if err != nil {
		return models.Buyer{}, errorbase.ErrDatabaseOperationFailed
	}

	return entity, nil
}

func (buyer *BuyerRepository) Delete(id int) error {

	if exist := buyer.buyerExist(id); !exist {
		return errorbase.ErrNotFound
	}

	_, err := buyer.db.Exec(querys.DeleteBuyer, id)
	if err != nil {
		return errorbase.ErrDatabaseOperationFailed
	}

	return nil
}

func (*BuyerRepository) validatemodels(entity models.Buyer) (map[string]any, error) {
	cardNumber, err := value_objects.NewCardNumberId(entity.Card_Number_Id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	firstName, err := value_objects.NewFirstName(entity.First_Name)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	lastName, err := value_objects.NewLastName(entity.Last_Name)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return map[string]any{
		"CardNumber": cardNumber,
		"FirstName":  firstName,
		"LastName":   lastName,
	}, nil
}

func (buyer *BuyerRepository) buyerExist(id int) bool {
	var exist bool
	err := buyer.db.QueryRow(querys.ExistsBuyer, id).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}

func (buyer *BuyerRepository) cardIdExist(id int, cardID string) bool {
	var exist bool
	if id > 0 {

		row := buyer.db.QueryRow(querys.SameBuyer, id)
		var buyer_model models.Buyer
		row.Scan(&buyer_model.Card_Number_Id)

		if cardID != buyer_model.Card_Number_Id {
			err := buyer.db.QueryRow(querys.ExistCardID, cardID).Scan(&exist)
			if err != nil {
				return false
			}
		} else {
			return false
		}

	} else {
		err := buyer.db.QueryRow(querys.ExistCardID, cardID).Scan(&exist)
		if err != nil {
			return false
		}
	}
	return exist
}

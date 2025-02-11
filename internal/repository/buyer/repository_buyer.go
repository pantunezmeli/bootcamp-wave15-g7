package buyer

import (
	"errors"

	//"dario.cat/mergo"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	buyerstorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/buyer_storage"

	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

type BuyerRepository struct {
	storage buyerstorage.IBuyerLoader
}

func NewBuyerRepository(storage buyerstorage.IBuyerLoader) *BuyerRepository {
	return &BuyerRepository{storage: storage}
}

func (buyer *BuyerRepository) GetAll() (map[int]models.Buyer, error) {

	buyers, err := buyer.storage.Load()
	if err != nil {
		return nil, errorbase.ErrStorageOperationFailed
	}
	list := make(map[int]models.Buyer)
	for key, value := range buyers {
		list[key] = value
	}

	return list, nil
}

func (buyer *BuyerRepository) GetById(id int) (models.Buyer, error) {
	buyers, err := buyer.storage.Load()
	if err != nil {
		return models.Buyer{}, errorbase.ErrStorageOperationFailed
	}

	byer_founded, ok := buyers[id]
	if !ok {
		return models.Buyer{}, errorbase.ErrNotFound
	}
	return byer_founded, nil
}

func (buyer *BuyerRepository) Create(entity models.Buyer) (models.Buyer, error) {
	buyers, err := buyer.storage.Load()
	if err != nil {
		return models.Buyer{}, errorbase.ErrStorageOperationFailed
	}

	if exist := searchCardId(buyers, entity.Card_Number_Id); exist {
		return models.Buyer{}, errorbase.ErrConflict
	}

	attributes, err := buyer.Validatemodels(entity)

	if err != nil {
		return models.Buyer{}, errorbase.ErrEmptyParameters
	}

	entity.Card_Number_Id = attributes["CardNumber"].(value_objects.CardNumberId).GetCardNumberId()
	entity.First_Name = attributes["FirstName"].(value_objects.FirstName).GetFirstName()
	entity.Last_Name = attributes["LastName"].(value_objects.LastName).GetLastName()

	entity.Id = getLastId(buyers)
	buyers[entity.Id] = entity

	err2 := buyer.storage.Save(entity)
	if err2 != nil {
		return models.Buyer{}, errorbase.ErrStorageOperationFailed
	}

	return entity, nil
}

func searchCardId(buyers map[int]models.Buyer, cardID int) bool {
	var found bool = false

	for _, value := range buyers {
		if value.Card_Number_Id == cardID {
			found = true
			break
		}
	}
	return found
}

func (buyer *BuyerRepository) Update(id int, entity models.Buyer) (models.Buyer, error) {
	buyers, err := buyer.storage.Load()
	if err != nil {
		return models.Buyer{}, errorbase.ErrStorageOperationFailed
	}

	element, ok := buyers[id]
	if !ok {
		return models.Buyer{}, errorbase.ErrNotFound
	}

	if entity.Card_Number_Id > 0 {
		var coincidence bool = false
		var i int = 0
		for i <= len(buyers) && !coincidence {
			coincidence = buyers[i].Card_Number_Id == entity.Card_Number_Id
			i++
		}

		if coincidence {
			return models.Buyer{}, errorbase.ErrConflict
		}
	}

	cardNumberId, err := value_objects.NewCardNumberId(entity.Card_Number_Id)
	if err == nil {
		element.Card_Number_Id = cardNumberId.GetCardNumberId()
	}

	firstName, err := value_objects.NewFirstName(entity.First_Name)
	if err == nil {
		element.First_Name = firstName.GetFirstName()
	}

	lastName, err := value_objects.NewLastName(entity.Last_Name)
	if err == nil {
		element.Last_Name = lastName.GetLastName()
	}

	buyers[id] = element
	err2 := buyer.storage.Save(element)
	if err2 != nil {
		return models.Buyer{}, errorbase.ErrStorageOperationFailed
	}

	return element, nil
}

func (buyer *BuyerRepository) Delete(id int) error {
	buyers, err := buyer.storage.Load()
	if err != nil {
		return errorbase.ErrStorageOperationFailed
	}

	buyerDelete, err := buyer.GetById(id)

	if err != nil {
		return err
	}
	delete(buyers, buyerDelete.Id)
	err2 := buyer.storage.Delete(id)
	if err2 != nil {
		return errorbase.ErrStorageOperationFailed
	}
	return nil
}

func getLastId(buyer map[int]models.Buyer) int {
	maxId := 0
	for id := range buyer {
		if id > maxId {
			maxId = id
		}
	}
	return maxId + 1
}

func (*BuyerRepository) Validatemodels(entity models.Buyer) (map[string]any, error) {
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

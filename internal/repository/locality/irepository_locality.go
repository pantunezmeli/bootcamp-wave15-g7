package locality

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

var (
	ErrProvinceNotFound = errors.New("province with id not found")
	ErrConnection = errors.New("connection error")
)

type LocalityRepository interface {
	Save(models.Locality) (modelSaved models.Locality, err error)
}
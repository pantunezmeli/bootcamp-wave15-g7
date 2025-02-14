package carrier

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

var (
	ErrInsertingData       = errors.New("error inserting data on database")
	ErrGettingLastID       = errors.New("error getting last insert ID")
	ErrConvertingID        = errors.New("error parsing ID")
	ErrForeignKeyViolation = errors.New("error invalid foreign key")
	ErrDuplicateEntry      = errors.New("error cid already exists")
	ErrDBGenericError      = errors.New("error database failed")
	ErrLocalityNotFound    = errors.New("locality does not exists")
)

type CarrierRepository struct {
	db *sql.DB
}

func NewCarrierRepository(db *sql.DB) *CarrierRepository {
	return &CarrierRepository{
		db: db,
	}
}

type CarrierByLocality struct {
	LocalityID    int
	LocalityName  string
	CarriesAmount int
}

// ! 1)
func (r *CarrierRepository) AddCarrierToDB(carrier models.Carrier) (models.Carrier, error) {
	query := `INSERT INTO carriers (cid, company_name, address, telephone, locality_id) VALUES  (?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query,
		carrier.Cid.GetCarrierCid(),
		carrier.CompanyName.GetCarrierCompanyName(),
		carrier.Address.GetAddress(),
		carrier.Telephone.GetTelephone(),
		carrier.LocalityId.GetLocalityId(),
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1452:
				return models.Carrier{}, ErrForeignKeyViolation
			case 1062:
				return models.Carrier{}, ErrDuplicateEntry
			default:
				return models.Carrier{}, ErrDBGenericError
			}
		}
		return models.Carrier{}, ErrInsertingData
	}

	lasInsertID, err := result.LastInsertId()
	if err != nil {
		return models.Carrier{}, ErrGettingLastID
	}

	newIdObj, err := value_objects.NewId(int(lasInsertID))
	if err != nil {
		return models.Carrier{}, ErrConvertingID
	}

	carrier.Id = newIdObj

	return carrier, nil

}

// ! 2)
func (r *CarrierRepository) GetCarriesAmountByLocalityID(id int) (CarrierByLocality, error) {
	query := `SELECT COALESCE(COUNT(c.id),0),l.locality_name, l.id
			FROM carriers c 
			RIGHT JOIN localities l ON c.locality_id = l.id 
			WHERE l.id = ?
			GROUP BY l.id, l.locality_name`

	var result CarrierByLocality

	err := r.db.QueryRow(query, id).Scan(
		&result.CarriesAmount,
		&result.LocalityName,
		&result.LocalityID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return CarrierByLocality{}, ErrLocalityNotFound
		}
		return CarrierByLocality{}, err
	}
	return result, nil

}

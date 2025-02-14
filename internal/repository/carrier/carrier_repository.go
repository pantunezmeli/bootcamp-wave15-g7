package carrier

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
	customErrors "github.com/pantunezmeli/bootcamp-wave15-g7/internal/errors"
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
				return models.Carrier{}, customErrors.ErrForeignKeyViolation
			case 1062:
				return models.Carrier{}, customErrors.ErrDuplicateEntry
			default:
				return models.Carrier{}, customErrors.ErrDBGenericError
			}
		}
		return models.Carrier{}, customErrors.ErrInsertingData
	}

	lasInsertID, err := result.LastInsertId()
	if err != nil {
		return models.Carrier{}, customErrors.ErrGettingLastID
	}

	newIdObj, err := value_objects.NewId(int(lasInsertID))
	if err != nil {
		return models.Carrier{}, customErrors.ErrConvertingID
	}

	carrier.Id = newIdObj

	return carrier, nil

}

// ! 2)
func (r *CarrierRepository) GetCarriesAmountByLocalityID(id *int) ([]CarrierByLocality, error) {
	query := `	SELECT 
					l.id AS locality_id,
					l.locality_name,
					COALESCE(COUNT(c.id), 0) AS carrier_count
				FROM localities l
				LEFT JOIN carriers c ON c.locality_id = l.id
				WHERE (? IS NULL OR l.id = ?)
				GROUP BY l.id, l.locality_name;`

	rows, err := r.db.Query(query, id, id)
	if err != nil {
		return nil, customErrors.ErrDBGenericError
	}
	defer rows.Close()

	var results []CarrierByLocality

	for rows.Next() {
		var result CarrierByLocality
		err := rows.Scan(&result.LocalityID, &result.LocalityName, &result.CarriesAmount)
		if err != nil {
			return nil, customErrors.ErrMappingData
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		return nil, customErrors.ErrLocalityNotFound
	}

	return results, nil
}

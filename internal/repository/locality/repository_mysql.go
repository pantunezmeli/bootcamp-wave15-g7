package locality

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	locality_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/locality"
)

type LocalityMySql struct {
	db *sql.DB
}

func NewLocalityMySql(db *sql.DB) *LocalityMySql {
	return &LocalityMySql{db}
}

func(r *LocalityMySql) Save(modelToSave models.Locality) (modelSaved models.Locality, err error) {
	result, err := r.db.Exec(`
	INSERT INTO localities (locality_name, province_id) VALUES (?, ?)`, modelToSave.Name, modelToSave.ProvinceId)
	if err != nil {
		var sqlError *mysql.MySQLError
		if errors.As(err, &sqlError){
			switch sqlError.Number {
			case 1452:
				err = ErrProvinceNotFound
			default:
				err = ErrConnection
			}
			return
		}
		err = ErrConnection
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		err = ErrConnection
		return
	}

	modelSaved = modelToSave
	modelSaved.Id = locality_vo.LocalityId(id)
	return
}

func (r *LocalityMySql) GetById(id int) (locality models.Locality, err error) {
	row := r.db.QueryRow(
		`
		SELECT id, locality_name, province_id
		from localities
		where id = ?
		`, id)
	if err = row.Err(); err != nil {
		err = ErrConnection
		return
	}

	if err = row.Scan(&locality.Id, &locality.Name, &locality.ProvinceId); err != nil{
		if errors.Is(err, sql.ErrNoRows){
			err = ErrLocalityNotFound
			return
		}
		err = ErrConnection
	}
	return
}
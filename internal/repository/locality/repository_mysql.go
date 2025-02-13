package locality

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	locality_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/locality"
	locality_dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/locality"

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

func (r *LocalityMySql) GetReportSellers(id *int) (reports []locality_dto.SellerReport, err error) {
	var rows *sql.Rows
	if id != nil {
		rows, err = r.db.Query(`
		SELECT l.id, l.locality_name, count(s.id)
		FROM localities l LEFT JOIN sellers s ON s.locality_id=l.id
		WHERE l.id = ?
		GROUP BY l.id
		`, &id)
	} else {
		rows, err = r.db.Query(`
		SELECT l.id, l.locality_name, count(s.id)
		FROM localities l LEFT JOIN sellers s ON s.locality_id=l.id
		GROUP BY l.id
		`)
	}
	if err != nil {
		err = ErrConnection
		return
	}

	for rows.Next(){
		var report locality_dto.SellerReport

		if err = rows.Scan(&report.LocalityId, &report.LocalityName, &report.SellersCount); err != nil {
			err = ErrConnection
			return
		}

		reports = append(reports, report)
	}

	if err = rows.Err(); err != nil {
		err = ErrConnection
	}

	if len(reports) == 0 {
		err = ErrLocalityNotFound
		return
	}
	return

}
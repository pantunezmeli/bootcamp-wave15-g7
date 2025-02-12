package productbatch

import (
	"database/sql"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/product_batch"
)

type ProductBatchRepository struct {
	db *sql.DB // Contiene una base de datos
}

// Inyección de la base de datos
func NewProductBatchRepository(db *sql.DB) *ProductBatchRepository {
	return &ProductBatchRepository{
		db: db,
	}
}

func (r *ProductBatchRepository) Store(batch *models.ProductBatch) (dto.ProductBatchResponse, error) {
	// Se inserta el lote de productos en la base de datos
	query := "INSERT INTO product_batch (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minumum_temperature, product_id, section_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	// result, err := r.db.Exec(query,
	// 	warehouse.WareHouseCode.GetWareHouseCode(),
	// 	warehouse.Address.GetAddress(),
	// 	warehouse.Telephone.GetTelephone(),
	// 	warehouse.MinimunCapacity.GetMinimunCapacity(),
	// 	warehouse.MinimunTemperature.GetMinimunTemperature(),
	// )

	_, err := r.db.Exec(query, batch.BatchNumber, batch.CurrentQuantity, batch.CurrentTemperature,
		batch.DueDate, batch.InitialQuantity, batch.ManufacturingDate, batch.ManufacturingHour, batch.MinumumTemperature,
		batch.ProductID, batch.SectionID)

	// Si hay un error se retorna
	if err != nil {
		return dto.ProductBatchResponse{}, err
	}

	// Se retorna el lote de productos
	return dto.GenerateProductBatchModelToDto(*batch), nil
}

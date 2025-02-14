package productbatch

import (
	"database/sql"
	"log"

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

	log.Println("Llego al repo")
	log.Println(batch.DueDate)

	query := "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := r.db.Exec(query, batch.BatchNumber, batch.CurrentQuantity, batch.CurrentTemperature,
		batch.DueDate, batch.InitialQuantity, batch.ManufacturingDate, batch.ManufacturingHour, batch.MinimumTemperature,
		batch.ProductID, batch.SectionID)

	// Si hay un error se retorna
	if err != nil {
		return dto.ProductBatchResponse{}, err
	}

	// Se retorna el lote de productos
	return dto.GenerateProductBatchModelToDto(*batch), nil
}

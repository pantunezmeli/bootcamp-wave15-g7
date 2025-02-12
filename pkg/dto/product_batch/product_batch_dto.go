package dto

import (
	"sort"
	"time"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

// ProductBatch is a struct that represents a product batch
type ProductBatchResponse struct {
	Id                 int       `json:"id"`
	BatchNumber        string    `json:"batch_number"`
	CurrentQuantity    int       `json:"current_quantity"`
	CurrentTemperature float64   `json:"current_temperature"`
	DueDate            time.Time `json:"due_date"`
	InitialQuantity    int       `json:"initial_quantity"`
	ManufacturingDate  time.Time `json:"manufacturing_date"`
	ManufacturingHour  time.Time `json:"manufacturing_hour"`
	MinumumTemperature float64   `json:"minimum_temperature"`
	ProductID          int       `json:"product_id"`
	SectionID          int       `json:"section_id"`
}

func GenerateProductBatchesResponseList(ProductBatches map[int]models.ProductBatch) []ProductBatchResponse {
	var list []ProductBatchResponse
	for _, value := range ProductBatches {
		list = append(list, ProductBatchResponse{
			Id:                 value.Id.GetId(),
			BatchNumber:        value.BatchNumber.GetBatchNumber(),
			CurrentQuantity:    value.CurrentQuantity.GetCurrentQuantity(),
			CurrentTemperature: value.CurrentTemperature.GetCurrentTemperature(),
			DueDate:            value.DueDate.GetDueDate(),
			InitialQuantity:    value.InitialQuantity.GetInitialQuantity(),
			ManufacturingDate:  value.ManufacturingDate.GetManufacturingDate(),
			ManufacturingHour:  value.ManufacturingHour.GetManufacturingHour(),
			MinumumTemperature: value.MinumumTemperature.GetMinumumTemperature(),
			ProductID:          value.ProductID.GetId(),
			SectionID:          value.SectionID.GetId(),
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})
	return list
}

func GenerateProductBatchModelToDto(ProductBatch models.ProductBatch) ProductBatchResponse {
	ProductBatchResponse := ProductBatchResponse{
		Id:                 ProductBatch.Id.GetId(),
		BatchNumber:        ProductBatch.BatchNumber.GetBatchNumber(),
		CurrentQuantity:    ProductBatch.CurrentQuantity.GetCurrentQuantity(),
		CurrentTemperature: ProductBatch.CurrentTemperature.GetCurrentTemperature(),
		DueDate:            ProductBatch.DueDate.GetDueDate(),
		InitialQuantity:    ProductBatch.InitialQuantity.GetInitialQuantity(),
		ManufacturingDate:  ProductBatch.ManufacturingDate.GetManufacturingDate(),
		ManufacturingHour:  ProductBatch.ManufacturingHour.GetManufacturingHour(),
		MinumumTemperature: ProductBatch.MinumumTemperature.GetMinumumTemperature(),
		ProductID:          ProductBatch.ProductID.GetId(),
		SectionID:          ProductBatch.SectionID.GetId(),
	}
	return ProductBatchResponse
}

func GenerateProductBatchDtoToModel(ProductBatchResponse ProductBatchResponse) (models.ProductBatch, error) {
	id, err := value_objects.NewId(ProductBatchResponse.Id)
	if err != nil {
		// handle the error appropriately
		return models.ProductBatch{}, err
	}
	batchNumber, err := value_objects.NewBatchNumber(ProductBatchResponse.BatchNumber)
	if err != nil {
		// handle the error appropriately
		return models.ProductBatch{}, err
	}
	currentQuantity, err := value_objects.NewCurrentQuantity(ProductBatchResponse.CurrentQuantity)
	if err != nil {
		return models.ProductBatch{}, err
	}
	currentTemperature, err := value_objects.NewCurrentTemperature(ProductBatchResponse.CurrentTemperature)
	if err != nil {
		return models.ProductBatch{}, err
	}
	dueDate, err := value_objects.NewDueDate(ProductBatchResponse.DueDate)
	if err != nil {
		return models.ProductBatch{}, err
	}
	initialQuantity, err := value_objects.NewInitialQuantity(ProductBatchResponse.InitialQuantity)
	if err != nil {
		return models.ProductBatch{}, err
	}
	manufacturingDate, err := value_objects.NewManufacturingDate(ProductBatchResponse.ManufacturingDate)
	if err != nil {
		return models.ProductBatch{}, err
	}
	manufacturingHour, err := value_objects.NewManufacturingHour(ProductBatchResponse.ManufacturingHour)
	if err != nil {
		return models.ProductBatch{}, err
	}
	minumumTemperature, err := value_objects.NewMinumumTemperature(ProductBatchResponse.MinumumTemperature)
	if err != nil {
		return models.ProductBatch{}, err
	}
	productId, err := value_objects.NewId(ProductBatchResponse.ProductID)
	if err != nil {
		return models.ProductBatch{}, err
	}
	sectionId, err := value_objects.NewId(ProductBatchResponse.SectionID)
	if err != nil {
		return models.ProductBatch{}, err
	}
	return models.ProductBatch{
		Id:                 id,
		BatchNumber:        batchNumber,
		CurrentQuantity:    currentQuantity,
		CurrentTemperature: currentTemperature,
		DueDate:            dueDate,
		InitialQuantity:    initialQuantity,
		ManufacturingDate:  manufacturingDate,
		ManufacturingHour:  manufacturingHour,
		MinumumTemperature: minumumTemperature,
		ProductID:          productId,
		SectionID:          sectionId,
	}, nil
}

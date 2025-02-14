package dto

import (
	"errors"
	"sort"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

// ProductBatch is a struct that represents a product batch
type ProductBatchResponse struct {
	Id                 int     `json:"id"`
	BatchNumber        string  `json:"batch_number"`
	CurrentQuantity    int     `json:"current_quantity"`
	CurrentTemperature float64 `json:"current_temperature"`
	DueDate            string  `json:"due_date"`
	InitialQuantity    int     `json:"initial_quantity"`
	ManufacturingDate  string  `json:"manufacturing_date"`
	ManufacturingHour  string  `json:"manufacturing_hour"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	ProductID          int     `json:"product_id"`
	SectionID          int     `json:"section_id"`
}

func GenerateProductBatchesResponseList(ProductBatches map[int]models.ProductBatch) []ProductBatchResponse {
	list := make([]ProductBatchResponse, 0, len(ProductBatches))
	for _, v := range ProductBatches {
		list = append(list, GenerateProductBatchModelToDto(v))
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})
	return list
}

func GenerateProductBatchModelToDto(ProductBatch models.ProductBatch) ProductBatchResponse {

	dueDate := ProductBatch.DueDate.Format("2006-01-02")

	manuFacturingDate := ProductBatch.ManufacturingDate.Format("2006-01-02")

	manuFacturingHour := ProductBatch.ManufacturingHour.Format("15:04:05")

	return ProductBatchResponse{
		Id:                 int(ProductBatch.Id),
		BatchNumber:        string(ProductBatch.BatchNumber),
		CurrentQuantity:    int(ProductBatch.CurrentQuantity),
		CurrentTemperature: float64(ProductBatch.CurrentTemperature),
		DueDate:            dueDate,
		InitialQuantity:    int(ProductBatch.InitialQuantity),
		ManufacturingDate:  manuFacturingDate,
		ManufacturingHour:  manuFacturingHour,
		MinimumTemperature: float64(ProductBatch.MinimumTemperature),
		ProductID:          ProductBatch.ProductID.GetId(),
		SectionID:          int(ProductBatch.SectionID),
	}
}

func GenerateProductBatchDtoToModel(ProductBatchResponse ProductBatchResponse) (models.ProductBatch, error) {

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
	minimumTemperature, err := value_objects.NewMinimumTemperature(ProductBatchResponse.MinimumTemperature)
	if err != nil {
		return models.ProductBatch{}, err
	}
	productId, err := value_objects.NewId(ProductBatchResponse.ProductID)
	if err != nil {
		return models.ProductBatch{}, errors.New("invalid product id")
	}
	sectionId, err := value_objects.NewId(ProductBatchResponse.SectionID)
	if err != nil {
		return models.ProductBatch{}, errors.New("invalid section id")
	}

	return models.ProductBatch{
		BatchNumber:        batchNumber,
		CurrentQuantity:    currentQuantity,
		CurrentTemperature: currentTemperature,
		DueDate:            dueDate,
		InitialQuantity:    initialQuantity,
		ManufacturingDate:  manufacturingDate,
		ManufacturingHour:  manufacturingHour,
		MinimumTemperature: minimumTemperature,
		ProductID:          productId,
		SectionID:          sectionId,
	}, nil
}

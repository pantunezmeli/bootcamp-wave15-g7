package dto

import (
	"sort"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects"
)

// ProductBatch is a struct that represents a product batch
type ProductBatchResponse struct {
	Id                 int                              `json:"id"`
	BatchNumber        value_objects.BatchNumber        `json:"batch_number"`
	CurrentQuantity    value_objects.CurrentQuantity    `json:"current_quantity"`
	CurrentTemperature value_objects.CurrentTemperature `json:"current_temperature"`
	DueDate            value_objects.DueDate            `json:"due_date"`
	InitialQuantity    value_objects.InitialQuantity    `json:"initial_quantity"`
	ManufacturingDate  value_objects.ManufacturingDate  `json:"manufacturing_date"`
	ManufacturingHour  value_objects.ManufacturingHour  `json:"manufacturing_hour"`
	MinimumTemperature value_objects.MinimumTemperature `json:"minimum_temperature"`
	ProductID          int                              `json:"product_id"`
	SectionID          int                              `json:"section_id"`
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
	return ProductBatchResponse{
		Id:                 ProductBatch.Id.GetId(),
		BatchNumber:        ProductBatch.BatchNumber,
		CurrentQuantity:    ProductBatch.CurrentQuantity,
		CurrentTemperature: ProductBatch.CurrentTemperature,
		DueDate:            ProductBatch.DueDate,
		InitialQuantity:    ProductBatch.InitialQuantity,
		ManufacturingDate:  ProductBatch.ManufacturingDate,
		ManufacturingHour:  ProductBatch.ManufacturingHour,
		MinimumTemperature: ProductBatch.MinimumTemperature,
		ProductID:          ProductBatch.ProductID.GetId(),
		SectionID:          ProductBatch.SectionID.GetId(),
	}
}

func GenerateProductBatchDtoToModel(ProductBatchResponse ProductBatchResponse) (models.ProductBatch, error) {
	id, err := value_objects.NewId(ProductBatchResponse.Id)
	if err != nil {
		// handle the error appropriately
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
		BatchNumber:        ProductBatchResponse.BatchNumber,
		CurrentQuantity:    ProductBatchResponse.CurrentQuantity,
		CurrentTemperature: ProductBatchResponse.CurrentTemperature,
		DueDate:            ProductBatchResponse.DueDate,
		InitialQuantity:    ProductBatchResponse.InitialQuantity,
		ManufacturingDate:  ProductBatchResponse.ManufacturingDate,
		ManufacturingHour:  ProductBatchResponse.ManufacturingHour,
		MinimumTemperature: ProductBatchResponse.MinimumTemperature,
		ProductID:          productId,
		SectionID:          sectionId,
	}, nil
}

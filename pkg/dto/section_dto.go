package dto

import (
	"sort"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

type SectionResponse struct {
	Id                  int `json:"id"`
	Section_Number      int `json:"section_number"`
	Current_Temperature int `json:"current_temperature"`
	Minimum_Temperature int `json:"minimum_temperature"`
	Current_Capacity    int `json:"current_capacity"`
	Minimum_Capacity    int `json:"minimum_capacity"`
	Maximum_Capacity    int `json:"maximum_capacity"`
	Warehouse_Id        int `json:"warehouse_id"`
	Product_Type_Id     int `json:"product_type_id"`
}

func GenerateBuyerResponseList(sections map[int]models.Section) []SectionResponse {
	var list []SectionResponse
	for _, value := range sections {
		list = append(list, SectionResponse{
			Id:                  value.Id,
			Section_Number:      value.Section_Number,
			Current_Temperature: value.Current_Temperature,
			Minimum_Temperature: value.Minimum_Temperature,
			Current_Capacity:    value.Current_Capacity,
			Minimum_Capacity:    value.Minimum_Capacity,
			Maximum_Capacity:    value.Maximum_Capacity,
			Warehouse_Id:        value.Warehouse_Id,
			Product_Type_Id:     value.Product_Type_Id,
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})
	return list
}

func GenerateSectionResponse(section models.Section) SectionResponse {
	sectionResponse := SectionResponse{
		Id:                  section.Id,
		Section_Number:      section.Section_Number,
		Current_Temperature: section.Current_Temperature,
		Minimum_Temperature: section.Minimum_Temperature,
		Current_Capacity:    section.Current_Capacity,
		Minimum_Capacity:    section.Minimum_Capacity,
		Maximum_Capacity:    section.Maximum_Capacity,
		Warehouse_Id:        section.Warehouse_Id,
		Product_Type_Id:     section.Product_Type_Id,
	}
	return sectionResponse
}

func GenerateSectionRequest(section SectionResponse) models.Section {
	SectionResponse := models.Section{
		Id:                  section.Id,
		Section_Number:      section.Section_Number,
		Current_Temperature: section.Current_Temperature,
		Minimum_Temperature: section.Minimum_Temperature,
		Current_Capacity:    section.Current_Capacity,
		Minimum_Capacity:    section.Minimum_Capacity,
		Maximum_Capacity:    section.Maximum_Capacity,
		Warehouse_Id:        section.Warehouse_Id,
		Product_Type_Id:     section.Product_Type_Id,
	}
	return SectionResponse
}

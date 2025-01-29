package dto

import (
	"sort"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
)

type BuyerResponse struct {
	Id             int    `json:"id"`
	Card_Number_Id int    `json:"card_number_id"`
	First_Name     string `json:"first_name"`
	Last_Name      string `json:"last_name"`
}

func GenerateResponseList(buyers map[int]model.Buyer) []BuyerResponse {
	var list []BuyerResponse
	for _, value := range buyers {
		list = append(list, BuyerResponse{
			Id:             value.Card_Number_Id,
			Card_Number_Id: value.Card_Number_Id,
			First_Name:     value.First_Name,
			Last_Name:      value.Last_Name,
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})
	return list
}

func GenerateBuyerResponse(buyer model.Buyer) BuyerResponse {
	buyerResponse := BuyerResponse{
		Id:             buyer.Id,
		Card_Number_Id: buyer.Card_Number_Id,
		First_Name:     buyer.First_Name,
		Last_Name:      buyer.Last_Name,
	}
	return buyerResponse
}

func GenerateBuyerRequeste(buyer BuyerResponse) model.Buyer {
	buyerResponse := model.Buyer{
		Id: buyer.Id,
		PersonAtributes: model.PersonAtributes{
			Card_Number_Id: buyer.Card_Number_Id,
			First_Name:     buyer.First_Name,
			Last_Name:      buyer.Last_Name,
		},
	}
	return buyerResponse
}

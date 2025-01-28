package buyerhandler

import (
	"sort"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/model"
)

func (b *BuyerHandler) generateResponseList(buyers map[int]model.Buyer) []model.BuyerResponse {
	var list []model.BuyerResponse
	for _, value := range buyers {
		list = append(list, model.BuyerResponse{
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

func (*BuyerHandler) generateBuyerResponse(buyer model.Buyer) model.BuyerResponse {
	buyerResponse := model.BuyerResponse{
		Id:             buyer.Id,
		Card_Number_Id: buyer.Card_Number_Id,
		First_Name:     buyer.First_Name,
		Last_Name:      buyer.Last_Name,
	}
	return buyerResponse
}

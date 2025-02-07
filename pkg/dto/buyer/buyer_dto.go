package dto

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

type BuyerResponse struct {
	Id             int    `json:"id"`
	Card_Number_Id int    `json:"card_number_id"`
	First_Name     string `json:"first_name"`
	Last_Name      string `json:"last_name"`
}

type BuyerUpdate struct {
	CardNumberID *int    `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
}

func ValidateBuyerFields(buyer BuyerUpdate) error {
	if buyer.CardNumberID != nil && *buyer.CardNumberID == 0 {
		return errors.New("card_number_id is empty")
	}

	if buyer.FirstName != nil && strings.TrimSpace(*buyer.FirstName) == "" {
		return errors.New("first_name is empty")
	}

	if buyer.LastName != nil && strings.TrimSpace(*buyer.LastName) == "" {
		return errors.New("last_name is empty")
	}

	return nil
}

func ConvertBuyerResponseToUpdate(response BuyerResponse) BuyerUpdate {
	return BuyerUpdate{
		CardNumberID: &response.Card_Number_Id,
		FirstName:    &response.First_Name,
		LastName:     &response.Last_Name,
	}
}

func ConvertBuyerUpdateToResponse(buyerUpdate BuyerUpdate) BuyerResponse {
	var cardNumberID int
	var firstName string
	var lastName string

	if buyerUpdate.CardNumberID != nil {
		cardNumberID = *buyerUpdate.CardNumberID
	}
	if buyerUpdate.FirstName != nil {
		firstName = *buyerUpdate.FirstName
	}
	if buyerUpdate.LastName != nil {
		lastName = *buyerUpdate.LastName
	}

	return BuyerResponse{
		Card_Number_Id: cardNumberID,
		First_Name:     firstName,
		Last_Name:      lastName,
	}
}

func GenerateResponseList(buyers map[int]models.Buyer) []BuyerResponse {
	var list []BuyerResponse
	for _, value := range buyers {
		list = append(list, BuyerResponse{
			Id:             value.Id,
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

func GenerateBuyerResponse(buyer models.Buyer) BuyerResponse {
	buyerResponse := BuyerResponse{
		Id:             buyer.Id,
		Card_Number_Id: buyer.Card_Number_Id,
		First_Name:     buyer.First_Name,
		Last_Name:      buyer.Last_Name,
	}
	return buyerResponse
}

func GenerateBuyerRequestUpdate(id int, buyerUpdate BuyerUpdate, buyerExist BuyerResponse) models.Buyer {
	fmt.Println(*buyerUpdate.CardNumberID)
	buyerResponse := models.Buyer{
		Id: id,
		PersonAtributes: models.PersonAtributes{
			Card_Number_Id: *buyerUpdate.CardNumberID,
			First_Name:     getStringFieldOrDefault(buyerUpdate.FirstName, buyerExist.First_Name),
			Last_Name:      getStringFieldOrDefault(buyerUpdate.LastName, buyerExist.Last_Name),
		},
	}

	return buyerResponse
}

func getStringFieldOrDefault(updateField *string, defaultValue string) string {
	if updateField != nil && *updateField != "" {
		return *updateField
	}
	return defaultValue
}

package seller_storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	seller_vo "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/value_objects/seller"
	seller_dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/seller"
)




func NewSellerJSONFile(path string) *SellerJSONFile {
	return &SellerJSONFile{
		path: path,
	}
}

type SellerJSONFile struct {
	path string
}

func (l *SellerJSONFile) Load() (sellerMap map[int]models.Seller, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	var sellersJSON []seller_dto.SellerDoc
	err = json.NewDecoder(file).Decode(&sellersJSON)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sellerMap = make(map[int]models.Seller)
	for _, s := range sellersJSON {
		Id, err :=seller_vo.NewSellerId(int(s.ID))
		if err != nil{
			return make(map[int]models.Seller), err
		}
		Cid, err :=seller_vo.NewCid(string(s.Cid))
		if err != nil{
			return make(map[int]models.Seller), err
		}
		CompanyName, err :=seller_vo.NewCompanyName(string(s.CompanyName))
		if err != nil{
			return make(map[int]models.Seller), err
		}
		Address, err :=seller_vo.NewSellerAddress(string(s.Address))
		if err != nil{
			return make(map[int]models.Seller), err
		}
		Telephone, err :=seller_vo.NewSellerTelephone(string(s.Telephone))
		if err != nil{
			return make(map[int]models.Seller), err
		}
		sellerMap[int(s.ID)] = models.Seller{
			ID: Id,
			SellerAttributes: models.SellerAttributes{
				Cid: Cid,
				CompanyName: CompanyName,
				Address: Address,
				Telephone: Telephone,
			},
		}
	}
	return
}

func (l *SellerJSONFile) Save(bd map[int]models.Seller) (err error) {
    file, err := os.Create(l.path)
    if err != nil {
        return ErrSavingFile
    }
   defer file.Close()

    sellers := make([]seller_dto.SellerDoc, 0, len(bd))
    for _, seller := range bd {
        sellerParsed :=seller_dto.SellerDoc{
			ID: seller.ID,
			Cid: seller.Cid,
			CompanyName: seller.CompanyName,
			Address: seller.Address,
			Telephone: seller.Telephone,
		}
		sellers = append(sellers, sellerParsed)
    }

    sellersJSON, err := json.MarshalIndent(sellers, "", "  ") 
    if err != nil {
        return ErrParsingData
    }

    _, err = file.Write(sellersJSON)
    if err != nil {
        return ErrSavingFile
    }

    return nil
}
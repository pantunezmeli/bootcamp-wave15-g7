package loader

import (
	"encoding/json"
	"fmt"
	"os"

	d "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
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

	var sellersJSON []dto.SellerDoc
	err = json.NewDecoder(file).Decode(&sellersJSON)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sellerMap = make(map[int]models.Seller)
	for _, s := range sellersJSON {
		Id, err := d.NewId(*s.ID)
		if err != nil{
			return make(map[int]models.Seller), err
		}
		Cid, err := d.NewCid(*s.Cid)
		if err != nil{
			return make(map[int]models.Seller), err
		}
		CompanyName, err := d.NewCompanyName(*s.CompanyName)
		if err != nil{
			return make(map[int]models.Seller), err
		}
		Address, err := d.NewAddress(*s.Address)
		if err != nil{
			return make(map[int]models.Seller), err
		}
		Telephone, err := d.NewTelephone(*s.Telephone)
		if err != nil{
			return make(map[int]models.Seller), err
		}
		sellerMap[*s.ID] = models.Seller{
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

    sellers := make([]dto.SellerDoc, 0, len(bd))
    for _, seller := range bd {
        sellerParsed := dto.SellerDoc{
			ID: seller.ID.Value(),
			Cid: seller.Cid.Value(),
			CompanyName: seller.CompanyName.Value(),
			Address: seller.Address.Value(),
			Telephone: seller.Telephone.Value(),
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
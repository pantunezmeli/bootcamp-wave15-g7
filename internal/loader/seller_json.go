package loader
import ("github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
d "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain"
"os"
"encoding/json"
)


func NewSellerJSONFile(path string) *SellerJSONFile {
	return &SellerJSONFile{
		path: path,
	}
}

type SellerJSONFile struct {
	path string
}

func (l *SellerJSONFile) Load() (v map[int]models.Seller, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var sellersJSON []dto.SellerDoc
	err = json.NewDecoder(file).Decode(&sellersJSON)
	if err != nil {
		return
	}

	sellerMap := make(map[int]models.Seller)
	for _, s := range sellersJSON {
		Id, err := d.NewId(s.ID)
		if err != nil{
			return make(map[int]models.Seller), err
		}
		Cid, err := d.NewCid(s.Cid)
		if err != nil{
			return make(map[int]models.Seller), err
		}
		CompanyName, err := d.NewCompanyName(s.CompanyName)
		if err != nil{
			return make(map[int]models.Seller), err
		}
		Address, err := d.NewAddress(s.Address)
		if err != nil{
			return make(map[int]models.Seller), err
		}
		Telephone, err := d.NewTelephone(s.Telephone)
		if err != nil{
			return make(map[int]models.Seller), err
		}
		sellerMap[s.ID] = models.Seller{
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

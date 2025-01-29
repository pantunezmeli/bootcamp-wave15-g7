package section

import (
	"errors"

	//"github.com/imdario/mergo"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/loader"
)

// NewSectionMap is a function that returns a new instance of SectionMap
func NewSectionMap(db map[int]models.Section) *SectionMap {
	// default db
	defaultDb := make(map[int]models.Section)
	if db != nil {
		defaultDb = db
	}
	return &SectionMap{db: defaultDb}
}

// SectionMap is a struct that represents a Section repository
type SectionMap struct {
	// db is a map of Sections
	db     map[int]models.Section
	loader loader.SectionLoader
}

// FindAll is a method that returns a map of all Sections
func (r *SectionMap) FindAll() (v map[int]models.Section, err error) {
	v = make(map[int]models.Section)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// FindByID is a method that returns a Section by its ID
func (r *SectionMap) FindByID(id int) (v models.Section, err error) {
	v, exists := r.db[id]
	if !exists {
		return v, errors.New("section not found")
	}
	return
}

// Create is a method that creates a new Section
func (r *SectionMap) Create(v models.Section) (err error) {
	if _, exists := r.db[v.Id]; exists {
		return errors.New("section already exists")
	}
	r.db[v.Id] = v
	return
}

// Patch is a method that updates a Section by its ID
/*func (r *SectionMap) Update(id int, v models.Section) (models.Section, error) {
	element, ok := r.db[id]
	if !ok {
		return models.Section{}, errorbase.ErrConflict
	}

	if sectionNumber, err := domain.NewSectionNumber(v.Section_Number); err == nil {
		element.Section_Number = sectionNumber.GetSectionNumber()
	}

	if currentTemperature, err := domain.NewCurrentTemperature(v.Current_Temperature); err == nil {
		element.Current_Temperature = currentTemperature.GetCurrentTemperature()
	}

	if minimumTemperature, err := domain.NewMinimumTemperature(v.Minimum_Temperature); err == nil {
		element.Minimum_Temperature = minimumTemperature.GetMinimumTemperature()
	}

	if currentCapacity, err := domain.NewCurrentCapacity(v.Current_Capacity); err == nil {
		element.Current_Capacity = currentCapacity.GetCurrentCapacity()
	}

	if minimumCapacity, err := domain.NewMinimumCapacity(v.Minimum_Capacity); err == nil {
		element.Minimum_Capacity = minimumCapacity.GetMinimumCapacity()
	}

	if maximumCapacity, err := domain.NewMaximumCapacity(v.Maximum_Capacity); err == nil {
		element.Maximum_Capacity = maximumCapacity.GetMaximumCapacity()
	}

	if warehouseId, err := domain.NewWarehouseId(v.Warehouse_Id); err == nil {
		element.Warehouse_Id = warehouseId.GetWarehouseId()
	}

	if productTypeId, err := domain.NewProductTypeId(v.Product_Type_Id); err == nil {
		element.Product_Type_Id = productTypeId.GetProductTypeId()
	}

	// if err := mergo.Merge(&element, v, mergo.WithOverride); err != nil {
	// 	return models.Section{}, err
	// }

	loaderV := loader.NewSectionJSONFile("../docs/db/section_data.json")
	loaderV.Save(element)
	r.db[id] = element
	return element, nil
}*/

// Delete is a method that deletes a Section by its ID
func (r *SectionMap) Delete(id int) (err error) {
	if _, exists := r.db[id]; !exists {
		return errors.New("section not found")
	}
	delete(r.db, id)
	return
}

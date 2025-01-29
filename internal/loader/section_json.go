package loader

import (
	"encoding/json"
	"os"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

// NewSectionJSONFile is a function that returns a new instance of SectionJSONFile
func NewSectionJSONFile(path string) *SectionJSONFile {
	return &SectionJSONFile{
		path: path,
	}
}

// SectionJSONFile is a struct that implements the LoaderSection interface
type SectionJSONFile struct {
	// path is the path to the file that contains the sections in JSON format
	path string
}

// Load is a method that loads the sections
func (l *SectionJSONFile) Load() (v map[int]models.Section, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var sectionsJSON []models.SectionDoc
	err = json.NewDecoder(file).Decode(&sectionsJSON)
	if err != nil {
		return
	}

	// serialize sections
	v = make(map[int]models.Section)
	for _, vh := range sectionsJSON {
		v[vh.Id] = models.Section{
			Id:                  vh.Id,
			Section_number:      vh.Section_number,
			Current_temperature: vh.Current_temperature,
			Minimum_temperature: vh.Minimum_temperature,
			Current_capacity:    vh.Current_capacity,
			Minimum_capacity:    vh.Minimum_capacity,
			Maximim_capacity:    vh.Maximim_capacity,
			Warehouse_id:        vh.Warehouse_id,
			Product_type_id:     vh.Product_type_id,
		}
	}

	return
}

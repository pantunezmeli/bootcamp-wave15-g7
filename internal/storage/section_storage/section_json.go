package loader

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

// NewSectionJSONFile is a function that returns a new instance of SectionJSONFile
func NewSectionJSONFile(path string) *SectionJSONFile {
	return &SectionJSONFile{path: path}
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
	var sectionsJSON []models.Section
	err = json.NewDecoder(file).Decode(&sectionsJSON)
	if err != nil {
		return
	}

	// serialize sections
	v = make(map[int]models.Section)
	for _, vh := range sectionsJSON {
		v[vh.Id] = models.Section{
			Id:                  vh.Id,
			Section_Number:      vh.Section_Number,
			Current_Temperature: vh.Current_Temperature,
			Minimum_Temperature: vh.Minimum_Temperature,
			Current_Capacity:    vh.Current_Capacity,
			Minimum_Capacity:    vh.Minimum_Capacity,
			Maximum_Capacity:    vh.Maximum_Capacity,
			Warehouse_Id:        vh.Warehouse_Id,
			Product_Type_Id:     vh.Product_Type_Id,
		}
	}

	return
}

func (l *SectionJSONFile) Save(section models.Section) error {
	var sections []models.Section

	file, err := os.Open(l.path)
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		_ = decoder.Decode(&sections)
	} else if !os.IsNotExist(err) {
		return errors.New("error reading")
	}

	updated := false
	for i := range sections {
		if sections[i].Id == section.Id {
			sections[i] = section
			updated = true
			break
		}
	}

	if !updated {
		sections = append(sections, section)
	}

	file, err = os.Create(l.path)
	if err != nil {
		return errors.New("error writing")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(sections); err != nil {
		return errors.New("error encoding")
	}

	return nil
}

func (l *SectionJSONFile) Delete(buyerID int) error {
	var sections []models.Section

	file, err := os.Open(l.path)
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		_ = decoder.Decode(&sections)
	} else if !os.IsNotExist(err) {
		return errors.New("error reading")
	}

	var updatedSections []models.Section
	found := false

	for _, buyer := range sections {
		if buyer.Id != buyerID {
			updatedSections = append(updatedSections, buyer)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("buyer not found")
	}

	file, err = os.Create(l.path)
	if err != nil {
		return errors.New("error writing")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(updatedSections); err != nil {
		return errors.New("error encoding")
	}

	return nil
}

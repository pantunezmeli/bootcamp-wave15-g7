package section

import (
	"dario.cat/mergo"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	sectionstorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/section"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

// StRepository is a struct that represents a Section repository
type StRepository struct {
	// storage is anRepository of Sections
	storage sectionstorage.ISectionLoader
}

// NewStRepository is a function that returns a new instance of StRepository
func NewStRepository(storage sectionstorage.ISectionLoader) *StRepository {
	// default storage

	return &StRepository{storage: storage}
}

// FindAll is a method that returns anRepository of all Sections
func (r *StRepository) FindAll() (map[int]models.Section, error) {
	sections, err := r.storage.Load()
	if err != nil {
		return nil, err
	}

	list := make(map[int]models.Section)
	for key, value := range sections {
		list[key] = value
	}

	return list, nil
}

// FinstorageyID is a method that returns a Section by its ID
func (r *StRepository) FindByID(id int) (entity models.Section, err error) {
	data, err := r.storage.Load()
	if err != nil {
		return models.Section{}, errorbase.ErrStorageOperationFailed
	}

	entity, exists := data[id]
	if !exists {
		return models.Section{}, errorbase.ErrNotFound
	}
	return entity, nil
}

// Create is a method that creates a new Section
func (r *StRepository) Create(entity models.Section) (models.Section, error) {

	// Load all sections
	sections, err := r.storage.Load()
	if err != nil {
		return models.Section{}, errorbase.ErrStorageOperationFailed
	}

	// Check if section number already exists
	for _, section := range sections {
		if section.Section_Number == entity.Section_Number {
			return models.Section{}, errorbase.ErrConflict
		}
	}

	// Set new ID
	entity.Id = getLastId(sections)
	sections[entity.Id] = entity

	// Save new section
	if err := r.storage.Save(entity); err != nil {
		return models.Section{}, err
	}
	return entity, nil
}

// Patch is a method that updates a Section by its ID
func (r *StRepository) Update(id int, entity models.Section) (models.Section, error) {
	// Load all sections
	data, err := r.storage.Load()
	if err != nil {
		return models.Section{}, errorbase.ErrStorageOperationFailed
	}

	// Check if section exists
	element, ok := data[id]
	if !ok {
		return models.Section{}, errorbase.ErrNotFound
	}

	// Merge the new data with the existing data
	if err := mergo.Merge(&element, entity, mergo.WithOverride); err != nil {
		return models.Section{}, err
	}

	data[id] = element
	if err := r.storage.Save(data[id]); err != nil {
		return models.Section{}, errorbase.ErrStorageOperationFailed
	}
	return element, nil
}

// Delete is a method that deletes a Section by its ID
func (r *StRepository) Delete(id int) (err error) {
	sections, err := r.storage.Load()
	if err != nil {
		return errorbase.ErrStorageOperationFailed
	}

	sectionDelete, err := r.FindByID(id)
	if err != nil {
		return errorbase.ErrNotFound
	}

	delete(sections, sectionDelete.Id)
	err = r.storage.Delete(id)
	if err != nil {
		return errorbase.ErrStorageOperationFailed
	}
	return nil
}

func getLastId(section map[int]models.Section) int {
	maxId := 0
	for id := range section {
		if id > maxId {
			maxId = id
		}
	}
	return maxId + 1
}

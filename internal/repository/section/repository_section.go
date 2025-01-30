package section

import (
	"errors"

	"dario.cat/mergo"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	sectionstorage "github.com/pantunezmeli/bootcamp-wave15-g7/internal/storage/section"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/errorbase"
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
	sections, _ := r.storage.Load()
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
		return models.Section{}, errors.New("section not found")
	}

	entity, exists := data[id]
	if !exists {
		return models.Section{}, errors.New("section not found")
	}
	return entity, nil
}

// Create is a method that creates a new Section
func (r *StRepository) Create(entity models.Section) (err error) {
	data, err := r.storage.Load()
	if err != nil {
		return err
	}
	if _, exists := data[entity.Id]; exists {
		return errors.New("section already exists")
	}
	data[entity.Id] = entity
	if err := r.storage.Save(entity); err != nil {
		return err
	}
	return
}

// Patch is a method that updates a Section by its ID
func (r *StRepository) Update(id int, entity models.Section) (models.Section, error) {
	data, err := r.storage.Load()
	if err != nil {
		return models.Section{}, err
	}

	element, ok := data[id]
	if !ok {
		return models.Section{}, errorbase.ErrConflict
	}

	if err := mergo.Merge(&element, entity, mergo.WithOverride); err != nil {
		return models.Section{}, err
	}

	data[id] = element
	if err := r.storage.Save(data[id]); err != nil {
		return models.Section{}, err
	}
	return element, nil
}

// Delete is a method that deletes a Section by its ID
func (r *StRepository) Delete(id int) (err error) {
	sections, _ := r.storage.Load()

	sectionDelete, err := r.FindByID(id)

	if err != nil {
		return err
	}
	delete(sections, sectionDelete.Id)
	err2 := r.storage.Delete(id)
	if err2 != nil {
		return errorbase.ErrStorageOperationFailed
	}
	return nil
}

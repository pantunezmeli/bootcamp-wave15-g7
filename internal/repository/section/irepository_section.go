package section

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
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
	db map[int]models.Section
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
func (r *SectionMap) Patch(id int, v models.Section) (err error) {
	if _, exists := r.db[id]; !exists {
		return errors.New("section not found")
	}
	r.db[id] = v
	return
}

// Delete is a method that deletes a Section by its ID
func (r *SectionMap) Delete(id int) (err error) {
	if _, exists := r.db[id]; !exists {
		return errors.New("section not found")
	}
	delete(r.db, id)
	return
}

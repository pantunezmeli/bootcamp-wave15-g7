package section

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

// SectionRepository is an interface that represents a section repository
type ISectionRepository interface {
	// FindAll is a method that returns a map of all sections
	FindAll() (map[int]models.Section, error)
	// FindByID is a method that returns a section by its ID
	FindByID(id int) (models.Section, error)
	// Create is a method that creates a new section
	Create(v models.Section) (models.Section, error)
	// Update is a method that updates a section by its ID
	Update(id int, e models.Section) (models.Section, error)
	// Delete is a method that deletes a section by its ID
	Delete(id int) error
}

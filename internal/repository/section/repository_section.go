package section

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

// VehicleRepository is an interface that represents a section repository
type SectionRepository interface {
	// FindAll is a method that returns a map of all sections
	FindAll() (v map[int]models.Section, err error)
	// FindByID is a method that returns a section by its ID
	FindByID(id int) (v models.Section, err error)
	// Create is a method that creates a new section
	Create(v models.Section) (err error)
	// Patch is a method that updates a section by its ID
	Update(id int, e models.Section) (updatedSection models.Section, err error)
	// Delete is a method that deletes a section by its ID
	Delete(id int) (err error)
}

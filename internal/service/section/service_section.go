package section

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

// SectionService is an interface that represents a section service
type SectionServiceV2 interface {
	// FindAll is a method that returns a map of all Sections
	ListSections() (v map[int]models.Section, err error)
	// GetSection is a method that returns a Section by its ID
	GetSection(id int) (v models.Section, err error)
	// CreateSection is a method that creates a new Section
	CreateSection(v models.Section) (err error)
	// PatchSection is a method that updates a Section by its ID
	PatchSection(id int, v models.Section) (updatedSection models.Section, err error)
	// DeleteSection is a method that deletes a Section by its ID
	DeleteSection(id int) (err error)
}

package section

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

// SectionService is an interface that represents a section service
type ISectionService interface {
	// FindAll is a method that returns a map of all Sections
	ListSections() ([]dto.SectionResponse, error)
	// GetSection is a method that returns a Section by its ID
	GetSection(id int) (dto.SectionResponse, error)
	// CreateSection is a method that creates a new Section
	CreateSection(v models.Section) (dto.SectionResponse, error)
	// PatchSection is a method that updates a Section by its ID
	PatchSection(id int, entity dto.SectionResponse) (dto.SectionResponse, error)
	// DeleteSection is a method that deletes a Section by its ID
	DeleteSection(id int) error
}

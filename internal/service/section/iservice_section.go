package section

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/section"
)

// NewSectionDefault is a function that returns a new instance of SectionDefault
func NewSectionDefault(rp section.SectionRepository) *SectionDefault {
	return &SectionDefault{rp: rp}
}

// SectionDefault is a struct that represents the default service for sections
type SectionDefault struct {
	// rp is the repository that will be used by the service
	rp section.SectionRepository
}

// ListSections is a method that returns a map of all Sections
func (s *SectionDefault) ListSections() (v map[int]models.Section, err error) {
	return s.rp.FindAll()
}

// GetSection is a method that returns a Section by its ID
func (s *SectionDefault) GetSection(id int) (v models.Section, err error) {
	return s.rp.FindByID(id)
}

// CreateSection is a method that creates a new Section
func (s *SectionDefault) CreateSection(v models.Section) (err error) {
	return s.rp.Create(v)
}

// PatchSection is a method that updates a Section by its ID
func (s *SectionDefault) PatchSection(id int, v models.Section) (updatedSection models.Section, err error) {
	return s.rp.Update(id, v)
}

// DeleteSection is a method that deletes a Section by its ID
func (s *SectionDefault) DeleteSection(id int) (err error) {
	return s.rp.Delete(id)
}

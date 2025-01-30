package section

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/section"
)

// NewSectionService is a function that returns a new instance of SectionService
func NewSectionService(rp section.ISectionRepository) *SectionService {
	return &SectionService{repository: rp}
}

// SectionService is a struct that represents the SerSectionService service for sections
type SectionService struct {
	// rp is the repository that will be used by the service
	repository section.ISectionRepository
}

// ListSections is a method that returns a map of all Sections
func (s *SectionService) ListSections() (v map[int]models.Section, err error) {
	return s.repository.FindAll()
}

// GetSection is a method that returns a Section by its ID
func (s *SectionService) GetSection(id int) (v models.Section, err error) {
	return s.repository.FindByID(id)
}

// CreateSection is a method that creates a new Section
func (s *SectionService) CreateSection(v models.Section) (err error) {
	return s.repository.Create(v)
}

// PatchSection is a method that updates a Section by its ID
func (s *SectionService) PatchSection(id int, v models.Section) (updatedSection models.Section, err error) {
	return s.repository.Update(id, v)
}

// DeleteSection is a method that deletes a Section by its ID
func (s *SectionService) DeleteSection(id int) (err error) {
	return s.repository.Delete(id)
}

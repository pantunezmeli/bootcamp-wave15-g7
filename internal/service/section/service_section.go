package section

import (
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/section"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/errorbase"
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
func (s *SectionService) ListSections() ([]dto.SectionResponse, error) {
	list, err := s.repository.FindAll()
	if err != nil {
		return nil, errorbase.ErrEmptyList
	}
	sections := dto.GenerateResponseList(list)
	return sections, nil
}

// GetSection is a method that returns a Section by its ID
func (s *SectionService) GetSection(id int) (dto.SectionResponse, error) {
	if id <= 0 {
		return dto.SectionResponse{}, errorbase.ErrInvalidId
	}
	section, err := s.repository.FindByID(id)
	if err != nil {
		return dto.SectionResponse{}, errorbase.ErrNotFound
	}
	dtoSection := dto.GenerateSectionResponse(section)
	return dtoSection, err
}

// CreateSection is a method that creates a new Section
func (s *SectionService) CreateSection(section models.Section) (dto.SectionResponse, error) {
	sectionCreated, err := s.repository.Create(section)
	sectionResponse := dto.GenerateSectionResponse(sectionCreated)
	return sectionResponse, err
}

// PatchSection is a method that updates a Section by its ID
func (s *SectionService) PatchSection(id int, entity dto.SectionResponse) (dto.SectionResponse, error) {

	sectionReq := dto.GenerateSectionRequest(entity)

	section, err := s.repository.Update(id, sectionReq)

	sectionResponse := dto.GenerateSectionResponse(section)
	return sectionResponse, err
}

// DeleteSection is a method that deletes a Section by its ID
func (s *SectionService) DeleteSection(id int) (err error) {
	s.repository.Delete(id)
	return nil
}

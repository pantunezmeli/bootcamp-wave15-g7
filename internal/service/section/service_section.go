package section

import (
	"errors"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/repository/section"
	dto "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto/section"
	errorbase "github.com/pantunezmeli/bootcamp-wave15-g7/pkg/error_base"
)

// NewSectionService is a function that returns a new instance of SectionService
func NewSectionService(rp section.ISectionRepository) *SectionService {
	return &SectionService{repository: rp}
}

// SectionService is a struct that represents the SerSectionService service for sections
type SectionService struct {
	// repository is the repository that will be used by the service
	repository section.ISectionRepository
}

// ListSections is a method that returns a map of all Sections
func (s *SectionService) ListSections() ([]dto.SectionResponse, error) {
	list, err := s.repository.FindAll()
	if err != nil {
		return nil, errorbase.ErrEmptyList
	}

	// Generate the response and return it
	sections := dto.GenerateSectionsResponseList(list)
	return sections, nil
}

// GetSection is a method that returns a Section by its ID
func (s *SectionService) GetSection(id int) (dto.SectionResponse, error) {
	// Check if the ID is valid
	if id <= 0 {
		return dto.SectionResponse{}, errorbase.ErrInvalidId
	}

	// Get the Section by ID
	section, err := s.repository.FindByID(id)
	if err != nil {
		return dto.SectionResponse{}, err
	}

	// Generate the response and return it
	dtoSection := dto.GenerateSectionResponse(section)
	return dtoSection, err
}

// CreateSection is a method that creates a new Section
func (s *SectionService) CreateSection(section models.Section) (dto.SectionResponse, error) {

	// Validate all parameters
	err := s.ValidateAllParameters(dto.GenerateSectionResponse(section))
	if err != nil {
		return dto.SectionResponse{}, err
	}

	// Create Section
	sectionCreated, err := s.repository.Create(section)
	if err != nil {
		return dto.SectionResponse{}, err
	}

	// Generate the response and return it
	sectionResponse := dto.GenerateSectionResponse(sectionCreated)
	return sectionResponse, err
}

// PatchSection is a method that updates a Section by its ID
func (s *SectionService) PatchSection(id int, entity dto.SectionResponse) (dto.SectionResponse, error) {
	// Check if the ID is valid
	if id <= 0 {
		return dto.SectionResponse{}, errorbase.ErrInvalidId
	}

	sectionReq := dto.GenerateSectionRequest(entity)

	section, err := s.repository.Update(id, sectionReq)
	if err != nil {
		return dto.SectionResponse{}, err
	}

	// Generate the response and return it
	sectionResponse := dto.GenerateSectionResponse(section)
	return sectionResponse, err
}

// DeleteSection is a method that deletes a Section by its ID
func (s *SectionService) DeleteSection(id int) (err error) {
	// Check if the ID is valid
	if id <= 0 {
		return errorbase.ErrInvalidId
	}

	// Delete Section
	err = s.repository.Delete(id)
	return
}

func (service *SectionService) ValidateAllParameters(reqBody dto.SectionResponse) (err error) {
	if reqBody.Section_Number == 0 {
		err = errors.New("section number is required")
		return
	}
	if reqBody.Current_Temperature == 0 {
		err = errors.New("current temperature is required")
		return
	}
	if reqBody.Minimum_Temperature == 0 {
		err = errors.New("minimum temperature is required")
		return
	}
	if condition := reqBody.Current_Temperature <= reqBody.Minimum_Temperature; condition {
		err = errors.New("current temperature must be greater than minimum temperature")
		return
	}
	if reqBody.Current_Capacity == 0 {
		err = errors.New("current capacity is required")
		return
	}
	if reqBody.Minimum_Capacity == 0 {
		err = errors.New("minimum capacity is required")
		return
	}
	if reqBody.Maximum_Capacity == 0 {
		err = errors.New("maximum capacity is required")
		return
	}
	if reqBody.Minimum_Capacity >= reqBody.Maximum_Capacity {
		err = errors.New("minimum capacity must be less than maximum capacity")
		return
	}
	if reqBody.Current_Capacity <= reqBody.Minimum_Capacity {
		err = errors.New("current capacity must be greater than minimum capacity")
		return
	}
	if reqBody.Maximum_Capacity <= reqBody.Current_Capacity {
		err = errors.New("maximum capacity must be greater than current capacity")
		return
	}
	if reqBody.Warehouse_Id == 0 {
		err = errors.New("warehouse id is required")
		return
	}
	if reqBody.Product_Type_Id == 0 {
		err = errors.New("product type id is required")
		return
	}
	return nil
}

package loader

import "github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"

// SectionLoader is an interface that represents the loader for Sections
type ISectionLoader interface {
	// Load is a method that loads the Sections
	Load() (v map[int]models.Section, err error)
	Save(section models.Section) error
	Delete(id int) error
}

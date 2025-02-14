package section

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/models"
)

type SectionRepository struct {
	db *sql.DB // Contiene una base de datos
}

// Inyección de la base de datos
func NewSectionRepository(db *sql.DB) *SectionRepository {
	return &SectionRepository{
		db: db,
	}
}

// FindAll is a method that returns anRepository of all Sections
func (r *SectionRepository) FindAll() ([]models.Section, error) {

	log.Println("Llego al repo")

	query := "SELECT id, section_number, current_capacity, current_temperature, maximum_capacity, minimum_capacity, minimum_temperature, product_type_id, warehouse_id FROM sections"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Tiro error")
		return nil, err
	}
	defer rows.Close()

	log.Println("Llego al repo 2")

	var sections []models.Section

	log.Println("llego al repo 3")

	// Iteración sobre los resultados
	for rows.Next() {
		var section models.Section

		err := rows.Scan(
			&section.Id,
			&section.Section_Number,
			&section.Current_Capacity,
			&section.Current_Temperature,
			&section.Maximum_Capacity,
			&section.Minimum_Capacity,
			&section.Minimum_Temperature,
			&section.Product_Type_Id,
			&section.Warehouse_Id,
		)
		if err != nil {
			return nil, err
		}

		log.Println(section)

		sections = append(sections, section)

		log.Println(sections)
	}
	return sections, nil
}

// FinstorageyID is a method that returns a Section by its ID
func (r *SectionRepository) FindByID(id int) (entity models.Section, err error) {
	var section models.Section
	query := "SELECT id, section_number, current_capacity, current_temperature, maximum_capacity, minimum_capacity, minimum_temperature, product_type_id, warehouse_id FROM sections WHERE id = ?"
	row := r.db.QueryRow(query, id)

	err = row.Scan(
		&section.Id,
		&section.Section_Number,
		&section.Current_Capacity,
		&section.Current_Temperature,
		&section.Maximum_Capacity,
		&section.Minimum_Capacity,
		&section.Minimum_Temperature,
		&section.Product_Type_Id,
		&section.Warehouse_Id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Section{}, fmt.Errorf("section with id %d not found", id)
		}
		return models.Section{}, err
	}
	return section, nil
}

// Create is a method that creates a new Section
func (r *SectionRepository) Create(section models.Section) (models.Section, error) {
	query := "INSERT INTO sections (section_number, current_capacity, current_temperature, maximum_capacity, minimum_capacity, minimum_temperature, product_type_id, warehouse_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(query,
		section.Section_Number,
		section.Current_Capacity,
		section.Current_Temperature,
		section.Maximum_Capacity,
		section.Minimum_Capacity,
		section.Minimum_Temperature,
		section.Product_Type_Id,
		section.Warehouse_Id,
	)

	if err != nil {
		return models.Section{}, err // Err
	}

	// Obtener el último ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return models.Section{}, err // ErrGettingLastInsertedID
	}

	section.Id = int(lastInsertID)

	return section, nil
}

// Update is a method that updates a Section by its ID
func (r *SectionRepository) Update(id int, section models.Section) error {
	query := "UPDATE sections SET section_number = ?, current_capacity = ?, current_temperature = ?, maximum_capacity = ?, minimum_capacity = ?, minimum_temperature = ?, product_type_id = ?, warehouse_id = ? WHERE id = ?"

	_, err := r.db.Exec(query,
		section.Section_Number,
		section.Current_Capacity,
		section.Current_Temperature,
		section.Maximum_Capacity,
		section.Minimum_Capacity,
		section.Minimum_Temperature,
		section.Product_Type_Id,
		section.Warehouse_Id,
		section.Id,
	)
	if err != nil {
		return err // ErrExecutingDB
	}

	return nil
}

// Delete is a method that deletes a Section by its ID
func (r *SectionRepository) Delete(id int) (err error) {
	query := "DELETE FROM sections WHERE id = ?"

	_, err = r.db.Exec(query,
		id)
	if err != nil {
		return err
	}

	return nil
}

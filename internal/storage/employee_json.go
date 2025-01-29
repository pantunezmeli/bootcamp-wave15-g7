package storage

import (
	"encoding/json"
	"os"

	"github.com/pantunezmeli/bootcamp-wave15-g7/internal/domain/model"
	"github.com/pantunezmeli/bootcamp-wave15-g7/pkg/dto"
)

func NewEmployeeJSONFile(path string) *EmployeeJSONFile {
	return &EmployeeJSONFile{
		path: path,
	}
}

type EmployeeJSONFile struct {
	path string
}

func (l *EmployeeJSONFile) Load() (employees map[int]model.Employee, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var EmployeesJSON []dto.EmployeeDoc
	err = json.NewDecoder(file).Decode(&EmployeesJSON)
	if err != nil {
		return
	}

	employees = make(map[int]model.Employee)
	for _, em := range EmployeesJSON {
		employees[em.Id], err = dto.EmployeeDtoToModel(em)
	}

	return
}

func (l *EmployeeJSONFile) Save(employees map[int]model.Employee) error {
	EmployeeList := make([]dto.EmployeeDoc, 0, len(employees))
	for _, e := range employees {
		EmployeeList = append(EmployeeList, dto.EmployeeModelToDto(e))
	}

	// Convertir la lista en JSON
	data, err := json.MarshalIndent(EmployeeList, "", "  ")
	if err != nil {
		return err
	}

	// Escribir el JSON generado en el archivo
	err = os.WriteFile(l.path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
